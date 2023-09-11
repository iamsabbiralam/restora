package storage

import (
	"database/sql"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	Pass     string = "PASS"
	TOTP     string = "TOTP"
	PINCode  string = "CODE"
	SMS      string = "SMS"
	Recovery string = "RECOVERY"
	EMail    string = "EMAIL"
)

const PAGESIZE int = 10

var (
	// NotFound is returned when the requested resource does not exist.
	NotFound = status.Error(codes.NotFound, "not found")
	// Conflict is returned when trying to create the same resource twice.
	Conflict = status.Error(codes.AlreadyExists, "conflict")
	// UsernameExists is returned when the username already exists in storage.
	UsernameExists = errors.New("username already exists")
	// EmailExists is returned when signup email already exists in storage.
	EmailExists = errors.New("email already exists")
	// InvCodeExists is returned when invitation code already exists in storage.
	InvCodeExists = errors.New("invitation code already exists")
	// Triggers when request arguments are invalid
	InvalidArgument = status.Error(codes.InvalidArgument, "invalid arguments")
)

var ErrNotFound = errors.New("not found")

type ActiveStatus int32

const (
	_ ActiveStatus = iota
	Active
	Inactive
)

type (
	User struct {
		ID        string         `db:"id"`
		Username  string         `db:"username"`
		Email     string         `db:"email"`
		Password  string         `db:"password"`
		Image     string         `db:"image"`
		Status    int32          `db:"status"`
		IsMFA     bool           `db:"is_mfa"`
		MFAType   string         `db:"mfa_type"`
		CreatedAt time.Time      `db:"created_at"`
		CreatedBy string         `db:"created_by"`
		UpdatedAt time.Time      `db:"updated_at"`
		UpdatedBy sql.NullString `db:"updated_by"`
		DeletedAt sql.NullTime   `db:"deleted_at,omitempty"`
		DeletedBy sql.NullString `db:"deleted_by,omitempty"`
		RoleNames sql.NullString `db:"role_names"`
		Count     int32
		RoleID    []string
	}

	UserInformation struct {
		ID        string         `db:"id"`
		UserID    string         `db:"user_id"`
		Username  string         `db:"username"`
		Image     string         `db:"image"`
		FirstName string         `db:"first_name"`
		LastName  string         `db:"last_name"`
		Email     string         `db:"email"`
		Mobile    string         `db:"mobile"`
		Gender    int            `db:"gender"`
		DOB       time.Time      `db:"dob"`
		Address   string         `db:"address"`
		City      string         `db:"city"`
		Country   string         `db:"country"`
		CreatedAt time.Time      `db:"created_at"`
		CreatedBy string         `db:"created_by"`
		UpdatedAt time.Time      `db:"updated_at"`
		UpdatedBy sql.NullString `db:"updated_by"`
		DeletedAt sql.NullTime   `db:"deleted_at,omitempty"`
		DeletedBy sql.NullString `db:"deleted_by,omitempty"`
	}

	FilterUser struct {
		SearchTerm   string
		Limit        int32
		Offset       int32
		SortBy       string
		SortByColumn string
		Status       ActiveStatus
		StartDate    string
		EndDate      string
	}

	Credential struct {
		ID            string       `db:"id"`
		Username      string       `db:"username"`
		Password      string       `db:"password"`
		Deleted       sql.NullTime `db:"deleted"`
		LastLogin     sql.NullTime `db:"last_login"`
		LastFailed    sql.NullTime `db:"last_failed"`
		ResetRequired sql.NullTime `db:"reset_required"`
		FailCount     int          `db:"fail_count"`
		EmailVerified bool         `db:"email_verified"`
	}

	Role struct {
		ID          string         `db:"id"`
		RoleName    string         `db:"name"`
		Description string         `db:"description"`
		Access      string         `db:"access"`
		Status      int32          `db:"status"`
		CreatedAt   time.Time      `db:"created_at"`
		CreatedBy   string         `db:"created_by"`
		UpdatedAt   time.Time      `db:"updated_at"`
		UpdatedBy   string         `db:"updated_by"`
		Delete      sql.NullTime   `db:"deleted_at,omitempty"`
		DeleteUID   sql.NullString `db:"deleted_by,omitempty"`
		Count       int
	}

	ListRoleFilter struct {
		SortBy     string `db:"sort_by"`
		SearchTerm string `db:"search_term"`
		Limit      int32  `db:"limit"`
		Offset     int32  `db:"offset"`
		Status     int32  `db:"status"`
	}

	ResAct struct {
		Resource string
		Action   string
		Public   bool
	}
)
