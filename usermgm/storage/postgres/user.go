package postgres

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/iamsabbiralam/restora/utility/random"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/argon2"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

const saltLen = 32

func (s *Storage) hashPassword(password, salt string) (string, error) {
	if password == "" {
		pw, err := random.String(20)
		if err != nil {
			return "", err
		}

		password = pw
	}

	if salt == "" {
		st, err := random.String(saltLen)
		if err != nil {
			return "", err
		}
		salt = st
	}

	pwHash := argon2.IDKey([]byte(password), []byte(salt), 5, 16*1024, 4, 32)
	return salt + hex.EncodeToString(pwHash), nil
}

const insertUser = `
	INSERT INTO users (
		username,
		email,
		password,
		status
	) VALUES (
		:username,
		:email,
		:password,
		:status
	) RETURNING
		id
`

func (s *Storage) CreateUser(ctx context.Context, user storage.User) (string, error) {
	stmt, err := s.db.PrepareNamed(insertUser)
	if err != nil {
		return "", err
	}

	var id string
	if err := stmt.Get(&id, user); err != nil {
		return "", err
	}

	return id, nil
}

const getUserByID = `
	SELECT
		*
	FROM
		users
	WHERE
		id = $1
	AND
		deleted_at IS NULL
`

func (s Storage) GetUserByID(ctx context.Context, id string) (*storage.User, error) {
	var u storage.User
	err := s.db.Get(&u, getUserByID, id)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

const getUserByUsername = `
SELECT
	*
FROM
	users
WHERE
	username = :username
AND
	deleted_at IS NULL
`

func (s Storage) GetUserByUsername(ctx context.Context, username string) (*storage.User, error) {
	var u storage.User
	stmt, err := s.db.PrepareNamed(getUserByUsername)
	if err != nil {
		return nil, err
	}

	arg := map[string]interface{}{
		"username": username,
	}

	if err := stmt.Get(&u, arg); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &u, nil
}

const getUserByEmail = `
SELECT
	*
FROM
	users
WHERE
	email = :email
AND
	deleted_at IS NULL
`

func (s Storage) GetUserByEmail(ctx context.Context, email string) (*storage.User, error) {
	var u storage.User
	stmt, err := s.db.PrepareNamed(getUserByEmail)
	if err != nil {
		return nil, err
	}

	arg := map[string]interface{}{
		"email": email,
	}
	if err := stmt.Get(&u, arg); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &u, nil
}

const updateUser = `
	UPDATE
		users
	SET
		username = COALESCE(NULLIF(:username, ''), username),
		email = COALESCE(NULLIF(:email, ''), email),
		is_mfa = COALESCE(NULLIF(:is_mfa, false), is_mfa),
		mfa_type = COALESCE(NULLIF(:mfa_type, ''), mfa_type),
		password = COALESCE(NULLIF(:password, ''), password),
		status = COALESCE(NULLIF(:status, 0), status),
		updated_by = COALESCE(NULLIF(:updated_by, ''), updated_by),
		updated_at = now()
	WHERE
		id = :id
	AND
		deleted_at IS NULL
	RETURNING
		*
`

func (s *Storage) UpdateUser(ctx context.Context, u storage.User) (*storage.User, error) {
	stmt, err := s.db.PrepareNamed(updateUser)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var user storage.User
	if err := stmt.Get(&user, u); err != nil {
		return nil, err
	}

	return &user, nil
}

const listUserQuery = `
	WITH cnt AS (select count(*) as count FROM users WHERE deleted_at IS NULL)
	SELECT
		usr.*,
		cnt.count,
		string_agg(roles.name, ', ') as role_names
	FROM 
		users as usr LEFT JOIN cnt on true 
		LEFT JOIN user_role as usrole on usr.id = usrole.user_id
		LEFT JOIN roles on usrole.role_id = roles.id
	WHERE 
		usr.deleted_at IS NULL
`

func (s *Storage) ListUsers(ctx context.Context, f storage.FilterUser) ([]storage.User, error) {
	searchQ := listUserQuery
	inp := []interface{}{}
	if f.SearchTerm != "" {
		searchTerms := strings.Split(f.SearchTerm, " ")
		searchQL := []string{}
		for _, n := range searchTerms {
			searchQL = append(searchQL, " (usr.username ILIKE ?) OR (usr.email ILIKE ?) ")
			nm := fmt.Sprintf("%%%s%%", n)
			inp = append(inp, nm)
			inp = append(inp, nm)
		}
		searchQ += " AND (" + strings.Join(searchQL, " OR ") + ") "
	}

	if f.SortBy != "ASC" { // default to descending on empty or invalid input
		f.SortBy = "DESC"
	}

	if f.Status != 0 {
		searchQ += " AND usr.status = ?"
		inp = append(inp, f.Status)
	}

	if f.StartDate != "" {
		searchQ += " AND usr.created_at >= ?"
		inp = append(inp, f.StartDate)
	}

	if f.EndDate != "" {
		searchQ += " AND usr.created_at <= ?"
		inp = append(inp, f.EndDate)
	}

	searchQ += " GROUP BY usr.id, cnt.count ORDER BY usr.created_at " + f.SortBy

	if f.Limit > 0 {
		searchQ += " LIMIT ?"
		inp = append(inp, f.Limit)
	}

	if f.Offset > 0 {
		searchQ += " OFFSET ?"
		inp = append(inp, f.Offset)
	}

	fullQuery, args, err := sqlx.In(searchQ, inp...)
	if err != nil {
		return nil, err
	}

	var usrs []storage.User
	if err := s.db.Select(&usrs, s.db.Rebind(fullQuery), args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return usrs, nil
}

const softDeleteUser = `
	UPDATE
		users
	SET
		deleted_at = now(),
		deleted_by = :deleted_by
	WHERE 
		id = :id
	AND
		deleted_at IS NULL
	RETURNING
		id
`

func (s *Storage) DeleteUser(ctx context.Context, u storage.User) error {
	stmt, err := s.db.PrepareNamedContext(ctx, softDeleteUser)
	if err != nil {
		return err
	}

	defer stmt.Close()
	var user storage.User
	if err := stmt.Get(&user, u); err != nil {
		return err
	}

	return nil
}

const deleteUser = `DELETE FROM users where id = $1`

func (s Storage) deleteUserPermanently(ctx context.Context, id string) error {
	row, err := s.db.Exec(deleteUser, id)
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

const deleteAllUsers = `DELETE FROM users`

func (s Storage) deleteUsersPermanently(ctx context.Context) error {
	row, err := s.db.ExecContext(ctx, deleteAllUsers)
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
		s.logger.Error("Unable to delete users")
		return storage.NotFound
	}

	return nil
}
