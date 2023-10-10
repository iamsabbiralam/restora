package storage

import (
	"database/sql"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrNotFound = errors.New("not found")

type ActiveStatus int32

const (
	_ ActiveStatus = iota
	Active
	Inactive
)

var (
	// NotFound is returned when the requested resource does not exist.
	NotFound = status.Error(codes.NotFound, "not found")
	// Conflict is returned when trying to create the same resource twice.
	Conflict = status.Error(codes.AlreadyExists, "conflict")
	// Triggers when request arguments are invalid
	InvalidArgument = status.Error(codes.InvalidArgument, "invalid arguments")
)

type (
	CRUDTimeDate struct {
		CreatedAt time.Time      `db:"created_at,omitempty"`
		CreatedBy string         `db:"created_by"`
		UpdatedAt time.Time      `db:"updated_at,omitempty"`
		UpdatedBy string         `db:"updated_by,omitempty"`
		DeletedAt sql.NullTime   `db:"deleted_at,omitempty"`
		DeletedBy sql.NullString `db:"deleted_by,omitempty"`
	}

	Category struct {
		ID     string `db:"id"`
		Name   string `db:"name"`
		Status int32  `db:"status"`
		Count  int32
		CRUDTimeDate
	}

	ListCategoryFilter struct {
		SortBy       string
		SearchTerm   string
		Limit        int32
		Offset       int32
		Status       ActiveStatus
		SortByColumn string
		StartDate    string
		EndDate      string
	}

	Brand struct {
		ID     string `db:"id"`
		Name   string `db:"name"`
		Status int32  `db:"status"`
		Count  int32
		CRUDTimeDate
	}

	ListBrandFilter struct {
		SortBy       string
		SearchTerm   string
		Limit        int32
		Offset       int32
		Status       ActiveStatus
		SortByColumn string
		StartDate    string
		EndDate      string
	}

	ResAct struct {
		Resource string
		Action   string
		Public   bool
	}
)
