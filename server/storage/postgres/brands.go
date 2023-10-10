package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/iamsabbiralam/restora/server/storage"
	"github.com/jmoiron/sqlx"
)

const insertBrand = `
INSERT INTO brands (
	name,
	status,
	created_by,
	created_at
) VALUES (
	:name, 
	:status,
	:created_by,
	now()
) RETURNING
	id
`

func (s *Storage) CreateBrand(ctx context.Context, req storage.Brand) (string, error) {
	if err := s.CreateBrandValidation(ctx, req); err != nil {
		return "", storage.InvalidArgument
	}

	stmt, err := s.db.PrepareNamed(insertBrand)
	if err != nil {
		return "", err
	}

	var id string
	if err := stmt.Get(&id, req); err != nil {
		return "", err
	}

	return id, nil
}

const getBrandByID = `
SELECT
	id,
	name,
	status
FROM 
	brands
WHERE 
	id = :id
AND
	deleted_at IS NULL
`

func (s *Storage) GetBrandByID(ctx context.Context, id string) (*storage.Brand, error) {
	var bra storage.Brand
	stmt, err := s.db.PrepareNamed(getBrandByID)
	if err != nil {
		return nil, err
	}

	arg := map[string]interface{}{
		"id": id,
	}
	if err := stmt.Get(&bra, arg); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.NotFound
		}

		return nil, err
	}

	return &bra, nil
}

const updateBrand = `
UPDATE 
	brands
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

func (s *Storage) UpdateBrand(ctx context.Context, sc storage.Brand) (*storage.Brand, error) {
	stmt, err := s.db.PrepareNamed(updateBrand)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var bra storage.Brand
	if err := stmt.Get(&bra, sc); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &bra, nil
}

const listBrand = `
	WITH cnt AS (select count(*) as count FROM brands WHERE deleted_at IS NULL)
	SELECT
		id,
		name,
		status,
		created_at,
		updated_at,
		cnt.count
	FROM
		brands
	LEFT JOIN
		cnt on true
	WHERE
		deleted_at IS NULL`

func (s *Storage) ListBrand(ctx context.Context, req storage.ListBrandFilter) ([]storage.Brand, error) {
	searchQ := listBrand
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

	var brands []storage.Brand
	if err := s.db.Select(&brands, s.db.Rebind(fullQuery), args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.NotFound
		}

		return nil, err
	}

	return brands, nil
}

const deleteBrand = `
UPDATE
	brands
SET
	deleted_at = now(),
	deleted_by = :deleted_by
WHERE 
	id = :id
AND
	deleted_at IS NULL
`

func (s *Storage) DeleteBrand(ctx context.Context, id, deletedBy string) error {
	stmt, err := s.db.PrepareNamedContext(ctx, deleteBrand)
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

const deleteAllBrands = `DELETE FROM brands`

func (s Storage) DeleteBrandsPermanently(ctx context.Context) error {
	row, err := s.db.ExecContext(ctx, deleteAllBrands)
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
		s.logger.Error("Unable to delete brands")
		return storage.NotFound
	}

	return nil
}
