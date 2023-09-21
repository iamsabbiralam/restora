package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/iamsabbiralam/restora/server/storage"
)

const insertCategory = `
	INSERT INTO categories (
		name,
		status,
		created_by
	) VALUES (
		:name, 
		:status,
		:created_by
	) RETURNING
		id
`

func (s *Storage) CreateCategory(ctx context.Context, req storage.Category) (string, error) {
	if err := s.CreateCategoryValidation(ctx, req); err != nil {
		return "", storage.InvalidArgument
	}

	stmt, err := s.db.PrepareNamed(insertCategory)
	if err != nil {
		return "", err
	}

	var id string
	if err := stmt.Get(&id, req); err != nil {
		return "", err
	}

	return id, nil
}

const getCategoryByID = `
	SELECT
		id,
		name,
		status
	FROM 
		categories
	WHERE 
		id = :id
	AND
		deleted_at IS NULL
`

func (s *Storage) GetCategoryByID(ctx context.Context, id string) (*storage.Category, error) {
	var cat storage.Category
	stmt, err := s.db.PrepareNamed(getCategoryByID)
	if err != nil {
		return nil, err
	}

	arg := map[string]interface{}{
		"id": id,
	}
	if err := stmt.Get(&cat, arg); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.NotFound
		}

		return nil, err
	}

	return &cat, nil
}

const updateCategory = `
	UPDATE 
		categories
	SET
		name = COALESCE(NULLIF(:name, ''), name),
		status = COALESCE(NULLIF(:status, 0), status),
		updated_by = COALESCE(NULLIF(:updated_by, ''), updated_by),
		updated_at = now()
	WHERE
		id = :id
	RETURNING
		*
`

func (s *Storage) UpdateCategory(ctx context.Context, sc storage.Category) (*storage.Category, error) {
	stmt, err := s.db.PrepareNamed(updateCategory)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var cat storage.Category
	if err := stmt.Get(&cat, sc); err != nil {
		return nil, err
	}

	return &cat, nil
}

const listCategories = `
	WITH cnt AS (select count(*) as count FROM categories WHERE deleted_at IS NULL)
	SELECT
		id,
		name,
		status,
		created_at,
		updated_at,
		cnt.count
	FROM
		categories
	LEFT JOIN
		cnt on true
	WHERE
		deleted_at IS NULL`

func (s *Storage) ListCategories(ctx context.Context, req storage.ListCategoryFilter) ([]storage.Category, error) {
	searchQ := listCategories
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

	var categories []storage.Category
	if err := s.db.Select(&categories, s.db.Rebind(fullQuery), args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.NotFound
		}

		return nil, err
	}

	return categories, nil
}

const deleteCategory = `
	DELETE FROM 
		categories
	WHERE 
		id = :id
`

func (s *Storage) DeleteCategory(ctx context.Context, id string) error {
	stmt, err := s.db.PrepareNamedContext(ctx, deleteCategory)
	if err != nil {
		return err
	}

	defer stmt.Close()
	arg := map[string]interface{}{
		"id": id,
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

const deleteAllCategories = `DELETE FROM categories`

func (s Storage) deleteCategoriesPermanently(ctx context.Context) error {
	row, err := s.db.ExecContext(ctx, deleteAllCategories)
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
		s.logger.Error("Unable to delete categories")
		return storage.NotFound
	}

	return nil
}
