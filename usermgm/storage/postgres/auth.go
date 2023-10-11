package postgres

import (
	"context"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

const login = `
	SELECT
		*
	FROM
		users
	WHERE
		email = :email AND
		status = 1 AND
		deleted_at IS NULL
`

func (s *Storage) Login(ctx context.Context, user storage.User) (*storage.User, error) {
	stmt, err := s.db.PrepareNamed(login)
	if err != nil {
		return nil, err
	}

	var getUser storage.User
	if err := stmt.Get(&getUser, user); err != nil {
		return nil, err
	}

	return &getUser, nil
}
