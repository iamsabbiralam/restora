package postgres

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/iamsabbiralam/restora/usermgm/storage"
)

var testCasesLoginByUsernameOrEmail = []struct {
	name     string
	In       storage.User
	want     *storage.User
	wantErr  bool
	setup    func(*testing.T, *Storage)
	teardown func(*testing.T, *Storage, string)
}{
	{
		name: "valid login",
		In: storage.User{
			Username: "testuser",
		},
		want:     &storage.User{Username: "testuser", Email: "testuser@example.com", Password: "password"},
		setup:    insertTestUser,
		teardown: deleteTestUser,
	},
	{
		name: "nonexistent_user",
		In: storage.User{
			Username: "nonexistentuser",
		},
		wantErr: true,
	},
}

func TestLoginByUsernameOrEmail(t *testing.T) {
	s := newTestStorage(t)
	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.User{},
			"ID",
			"Password",
			"IsMFA",
			"MFAType",
			"CreatedAt",
			"CreatedBy",
			"UpdatedAt",
			"UpdatedBy",
			"DeletedAt",
			"DeletedBy",
		),
	}

	for _, tc := range testCasesLoginByUsernameOrEmail {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			got, err := s.Login(context.Background(), tc.In)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.TestLoginByUsernameOrEmail() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.TestLoginByUsernameOrEmail() diff = %v", cmp.Diff(got, tc.want, opts...))
			}
		})
	}
}
