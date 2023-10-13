package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/iamsabbiralam/restora/server/storage"
)

const insertRecipe = `
INSERT INTO recipes (
	title,
	ingredient,
	image,
	description,
	author_social_link,
	read_count,
	serving_amount,
	cooking_time,
	is_used,
	user_id,
	status,
	created_by,
	created_at
) VALUES (
	:title,
	:ingredient,
	:image,
	:description,
	:author_social_link,
	:read_count,
	:serving_amount,
	:cooking_time,
	:is_used,
	:user_id,
	:status,
	:created_by,
	now()
) RETURNING
	id
`

func (s *Storage) CreateRecipe(ctx context.Context, req storage.Recipe) (string, error) {
	if err := s.CreateRecipeValidation(ctx, req); err != nil {
		return "", storage.InvalidArgument
	}

	stmt, err := s.db.PrepareNamed(insertRecipe)
	if err != nil {
		return "", err
	}

	var id string
	if err := stmt.Get(&id, req); err != nil {
		return "", err
	}

	return id, nil
}

const getRecipeByID = `
	SELECT
		title,
		ingredient,
		image,
		description,
		author_social_link,
		read_count,
		serving_amount,
		cooking_time,
		is_used,
		user_id,
		status
	FROM 
		recipes
	WHERE 
		id = :id
	AND
		deleted_at IS NULL
`

func (s *Storage) GetRecipeByID(ctx context.Context, id string) (*storage.Recipe, error) {
	var sr storage.Recipe
	stmt, err := s.db.PrepareNamed(getRecipeByID)
	if err != nil {
		return nil, err
	}

	arg := map[string]interface{}{
		"id": id,
	}
	if err := stmt.Get(&sr, arg); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.NotFound
		}

		return nil, err
	}

	return &sr, nil
}

const updateRecipe = `
	UPDATE 
		recipes
	SET
		title = COALESCE(NULLIF(:title, ''), title),
		ingredient = COALESCE(NULLIF(:ingredient, ''), ingredient),
		image = COALESCE(NULLIF(:image, ''), image),
		description = COALESCE(NULLIF(:description, ''), description),
		author_social_link = COALESCE(NULLIF(:author_social_link, ''), author_social_link),
		read_count = COALESCE(NULLIF(:read_count, 0), read_count),
		serving_amount = COALESCE(NULLIF(:serving_amount, 0), serving_amount),
		cooking_time = COALESCE(NULLIF(:cooking_time, 0), cooking_time),
		is_used = COALESCE(NULLIF(:is_used, 0), is_used),
		user_id = COALESCE(NULLIF(:user_id, 0), user_id),
		status = COALESCE(NULLIF(:status, 0), status),
		updated_by = COALESCE(NULLIF(:updated_by, ''), updated_by),
		updated_at = now()
	WHERE
		id = :id
	RETURNING
		*
`

func (s *Storage) UpdateRecipe(ctx context.Context, sc storage.Recipe) (*storage.Recipe, error) {
	stmt, err := s.db.PrepareNamed(updateRecipe)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var sr storage.Recipe
	if err := stmt.Get(&sr, sc); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &sr, nil
}


const listRecipes = `
	WITH cnt AS (select count(*) as count FROM recipes WHERE deleted_at IS NULL)
	SELECT
		id,
		title,
		ingredient,
		image,
		description,
		author_social_link,
		read_count,
		serving_amount,
		cooking_time,
		is_used,
		user_id,
		status,
		created_at,
		updated_at,
		cnt.count
	FROM
		recipes
	LEFT JOIN
		cnt on true
	WHERE
		deleted_at IS NULL`

func (s *Storage) ListRecipe(ctx context.Context, req storage.ListRecipeFilter) ([]storage.Recipe, error) {
	searchQ := listRecipes
	inp := []interface{}{}
	if req.SearchTerm != "" {
		searchQL := []string{}
		searchQL = append(searchQL, " AND (name ILIKE ? ) ")
		nm := fmt.Sprintf("%%%s%%", req.SearchTerm)
		inp = append(inp, nm)
		searchQ += strings.Join(searchQL, " ")
	}

	if req.Status != 0 {
		searchQ += " AND status = ?"
		inp = append(inp, req.Status)
	}

	if req.StartDate != "" {
		searchQ += " AND created_at >= ?"
		inp = append(inp, req.StartDate)
	}

	if req.EndDate != "" {
		searchQ += " AND created_at <= ?"
		inp = append(inp, req.EndDate)
	}

	if req.SortBy != "ASC" {
		req.SortBy = "DESC"
	}

	sortByColumn := "created_at"
	if req.SortByColumn == "name" || req.SortByColumn == "status" || req.SortByColumn == "updated_at" {
		sortByColumn = req.SortByColumn
	}

	searchQ += " ORDER BY ?" + req.SortBy
	inp = append(inp, sortByColumn)

	if req.Limit > 0 {
		searchQ += " LIMIT ?"
		inp = append(inp, req.Limit)
	}

	if req.Offset > 0 {
		searchQ += " OFFSET ?"
		inp = append(inp, req.Offset)
	}

	fullQuery, args, err := sqlx.In(searchQ, inp...)
	if err != nil {
		return nil, err
	}

	var recipes []storage.Recipe
	if err := s.db.Select(&recipes, s.db.Rebind(fullQuery), args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.NotFound
		}

		return nil, err
	}

	return recipes, nil
}

const deleteRecipe = `
UPDATE
	recipes
SET
	deleted_at = now(),
	deleted_by = :deleted_by
WHERE 
	id = :id
AND
	deleted_at IS NULL
`

func (s *Storage) DeleteRecipe(ctx context.Context, id, deletedBy string) error {
	stmt, err := s.db.PrepareNamedContext(ctx, deleteRecipe)
	if err != nil {
		return err
	}

	defer stmt.Close()
	arg := map[string]interface{}{
		"id":         id,
		"deleted_by": deletedBy,
	}

	row, err := stmt.Exec(arg)
	if err != nil {
		return err
	}

	rowCount, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowCount <= 0 {
		return storage.NotFound
	}

	return nil
}

const deleteAllRecipes = `DELETE FROM recipes`

func (s Storage) DeleteAllRecipePermanently(ctx context.Context) error {
	row, err := s.db.ExecContext(ctx, deleteAllRecipes)
	if err != nil {
		s.logger.WithError(err)
		return err
	}

	rowCount, err := row.RowsAffected()
	if err != nil {
		s.logger.WithError(err)
		return err
	}

	if rowCount <= 0 {
		s.logger.Error("Unable to delete recipes")
		return storage.NotFound
	}

	return nil
}
