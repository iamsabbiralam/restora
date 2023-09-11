package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

const insertUserInformation = `
	INSERT INTO user_information (
		user_id
	) VALUES (
		:user_id
	) RETURNING
		id
`

func (s *Storage) CreateUserInformation(ctx context.Context, ui storage.UserInformation) (string, error) {
	if err := s.CreateUserInformationValidation(ctx, ui); err != nil {
		return "", storage.InvalidArgument
	}

	stmt, err := s.db.PrepareNamedContext(ctx, insertUserInformation)
	if err != nil {
		return "", err
	}

	var id string
	if err := stmt.Get(&id, ui); err != nil {
		return "", err
	}

	return id, nil
}

const getUserInformation = `
	SELECT
		userinfo.*
	FROM
		user_information as userinfo
    		LEFT JOIN users as usr on usr.id = userinfo.user_id
	WHERE
		user_id = 'b6ddbe32-3d7e-4828-b2d7-da9927846e6b'
	AND
		userinfo.deleted_at IS NULL
`

func (s *Storage) GetUserInformation(ctx context.Context, userID string) (*storage.UserInformation, error) {
	var profile storage.UserInformation
	stmt, err := s.db.PrepareNamedContext(ctx, getUserInformation)
	if err != nil {
		return nil, err
	}

	arg := map[string]interface{}{
		"user_id": userID,
	}
	if err := stmt.Get(&profile, arg); err != nil {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, storage.ErrNotFound
	}
	fmt.Println("---profile---")
	fmt.Println(profile)
	fmt.Println("---profile---")
	return &profile, nil
}

const updateUserInformation = `
	UPDATE
		user_information
	SET
		image = COALESCE(NULLIF(:image, ''), image),
		first_name = COALESCE(NULLIF(:first_name, ''), first_name),
		last_name = COALESCE(NULLIF(:last_name, ''), last_name),
		mobile = COALESCE(NULLIF(:mobile, ''), mobile),
		gender = COALESCE(NULLIF(:gender, 0), gender),
		dob = COALESCE(NULLIF(:dob, DATE '0001-01-01'), dob),
		address = COALESCE(NULLIF(:address, ''), address),
		city = COALESCE(NULLIF(:city, ''), city),
		country = COALESCE(NULLIF(:country, ''), country),
		updated_at = now(),
		updated_by = COALESCE(NULLIF(:updated_by, ''), updated_by)
	WHERE
		user_id = :user_id AND deleted_at IS NULL
	RETURNING
		*
`

func (s *Storage) UpdateUserInformation(ctx context.Context, ui storage.UserInformation) (*storage.UserInformation, error) {
	stmt, err := s.db.PrepareNamed(updateUserInformation)
	if err != nil {
		return &storage.UserInformation{}, err
	}

	defer stmt.Close()
	var profile storage.UserInformation
	if err := stmt.Get(&profile, ui); err != nil {
		return &storage.UserInformation{}, err
	}

	return &profile, nil
}

const deleteUserInformation = `
	UPDATE
		user_information
	SET
		deleted_at = now(),
		deleted_by = :deleted_by
	WHERE
		user_id = :user_id
`

func (s *Storage) DeleteUserInformation(ctx context.Context, userID, deletedBy string) error {
	if err := s.DeleteUserInformationValidation(ctx, userID); err != nil {
		return err
	}

	stmt, err := s.db.PrepareNamedContext(ctx, deleteUserInformation)
	if err != nil {
		return err
	}

	defer stmt.Close()
	arg := map[string]interface{}{
		"user_id":    userID,
		"deleted_by": deletedBy,
	}
	if _, err := stmt.Exec(arg); err != nil {
		return err
	}

	return nil
}

const deleteAllUserInformation = `DELETE FROM user_information`

func (s Storage) DeleteUserInformationPermanently(ctx context.Context) error {
	row, err := s.db.ExecContext(ctx, deleteAllUserInformation)
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
