package postgres

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

func insertTestUser(t *testing.T, s *Storage) {
	// insert test user into the "users" table
	got, err := s.CreateUser(context.TODO(), storage.User{Username: "testuser", Email: "testuser@example.com", Password: "password"})
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}

	if got == "" {
		t.Fatalf("Unable to create user: %v", err)
	}
}

func deleteTestUser(t *testing.T, s *Storage, id string) {
	// delete test user from the "users" table
	if err := s.deleteUserPermanently(context.TODO(), id); err != nil {
		t.Fatalf("Unable to delete user: %v", err)
	}
}

func deleteAllTestUser(t *testing.T, s *Storage) {
	// delete test user from the "users" table
	if err := s.deleteUsersPermanently(context.TODO()); err != nil {
		t.Fatalf("Unable to delete users: %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	s := newTestStorage(t)
	testCases := []struct {
		name     string
		in       storage.User
		want     string
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage)
	}{
		{
			name: "user creation success",
			in: storage.User{
				Username: "testuser",
				Email:    "testuser@example.com",
				Password: "password",
				Status:   1,
			},
			wantErr:  false,
			teardown: deleteAllTestUser,
		}, {
			name: "user creation failed",
			in: storage.User{
				Username: "testuser",
				Email:    "testuser@example.com",
				Password: "password",
				Status:   1,
			},
			wantErr:  true,
			setup:    insertTestUser,
			teardown: deleteAllTestUser,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			got, err := s.CreateUser(context.TODO(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.CreateUser() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && got == "" {
				t.Errorf("Storage.CreateUser() want uuid, got %v", got)
				return
			}

			if tc.teardown != nil {
				tc.teardown(t, s)
			}
		})
	}
}

var getUserByIDtestCases = []struct {
	name     string
	id       string
	want     *storage.User
	wantErr  bool
	teardown func(*testing.T, *Storage, string)
	setup    func(*testing.T, *Storage, string)
}{
	{
		name:     "validUserID",
		want:     &storage.User{Username: "testuser", Email: "testuser@example.com", Password: "password"},
		teardown: deleteTestUser,
	},
	{
		name:    "nonExistentUserID",
		id:      "nonexistentuserid",
		wantErr: true,
	},
}

func TestGetUserByID(t *testing.T) {
	s := newTestStorage(t)
	uid, err := s.CreateUser(context.TODO(), storage.User{Username: "testuser", Email: "testuser@example.com", Password: "password"})
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}

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

	for _, tc := range getUserByIDtestCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.id != "" {
				uid = tc.id
			}

			got, err := s.GetUserByID(context.Background(), uid)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.GetUserByID() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.GetUserByID() diff = %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, s, uid)
			}
		})
	}
}

var testCasesGetUserByUsername = []struct {
	name     string
	username string
	want     *storage.User
	wantErr  bool
	setup    func(*testing.T, *Storage)
	teardown func(*testing.T, *Storage, string)
}{
	{
		name:     "valid username",
		username: "testuser",
		want:     &storage.User{Username: "testuser", Email: "testuser@example.com", Password: "password"},
		setup:    insertTestUser,
		teardown: deleteTestUser,
	},
	{
		name:     "nonexistent user",
		username: "nonexistentuser",
		wantErr:  true,
	},
}

func TestGetUserByUsername(t *testing.T) {
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

	for _, tc := range testCasesGetUserByUsername {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			got, err := s.GetUserByUsername(context.Background(), tc.username)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.GetUserByUsername() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.GetUserByUsername() diff = %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, s, got.ID)
			}
		})
	}
}

var testCasesGetUserByEmail = []struct {
	name     string
	email    string
	want     *storage.User
	wantErr  bool
	setup    func(*testing.T, *Storage)
	teardown func(*testing.T, *Storage, string)
}{
	{
		name:     "valid username",
		email:    "testuser@example.com",
		want:     &storage.User{Username: "testuser", Email: "testuser@example.com", Password: "password"},
		setup:    insertTestUser,
		teardown: deleteTestUser,
	},
	{
		name:    "nonexistent user email",
		email:   "nonexistenttestuser@example.com",
		wantErr: true,
	},
}

func TestGetUserByEmail(t *testing.T) {
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

	for _, tc := range testCasesGetUserByEmail {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			got, err := s.GetUserByEmail(context.Background(), tc.email)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.GetUserByEmail() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.GetUserByEmail() diff = %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, s, got.ID)
			}
		})
	}
}

var testUpdateUser = []struct {
	name     string
	in       storage.User
	want     *storage.User
	wantErr  bool
	setup    func(*testing.T, *Storage)
	teardown func(*testing.T, *Storage, string)
}{
	{
		name: "update user success",
		in: storage.User{
			Username: "testuser",
			Status:   1,
			IsMFA:    true,
			MFAType:  "email",
		},
		want: &storage.User{
			Username: "testuser",
			Email:    "testuser@example.com",
			Password: "password",
			Status:   1,
			IsMFA:    true,
			MFAType:  "email",
		},
		setup:    insertTestUser,
		teardown: deleteTestUser,
	},
}

func TestUpdateUser(t *testing.T) {
	s := newTestStorage(t)
	for _, tc := range testUpdateUser {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}
			user, err := s.GetUserByUsername(context.TODO(), tc.in.Username)
			if err != nil {
				t.Fatal(err)
			}

			tc.in.ID = user.ID
			tc.in.UpdatedBy = sql.NullString{String: user.ID, Valid: true}
			got, err := s.UpdateUser(context.Background(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.UpdateUser() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(
					storage.User{},
					"ID",
					"Password",
					"CreatedAt",
					"CreatedBy",
					"UpdatedAt",
					"UpdatedBy",
					"DeletedAt",
					"DeletedBy",
				),
			}
			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.UpdateUser() diff = %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, s, user.ID)
			}
		})
	}
}

func TestListUser(t *testing.T) {
	ts := newTestStorage(t)
	testsList := []struct {
		name     string
		in       storage.FilterUser
		want     []storage.User
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage, string)
	}{
		{
			name: "GET_LIST_USER_SUCCESS",
			in: storage.FilterUser{
				SearchTerm: "user search",
			},
			want: []storage.User{
				{
					Email:    "testuser@example.com",
					Username: "testuser",
					Status:   0,
				},
			},
			teardown: deleteTestUser,
		},
	}

	opts := []cmp.Option{
		cmpopts.IgnoreFields(storage.User{}, "ID", "Count", "Password", "IsMFA", "MFAType", "CreatedAt", "DeletedAt", "UpdatedAt", "DeletedBy", "CreatedBy", "UpdatedBy"),
	}

	deleteAllTestUser(t, ts)
	insertTestUser(t, ts)

	for _, tt := range testsList {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.ListUsers(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(tt.want, got, opts...) {
				t.Error("(-want +got): ", cmp.Diff(tt.want, got, opts...))
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	s := newTestStorage(t)
	testCases := []struct {
		name     string
		in       storage.User
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage, string)
	}{
		{
			name: "user deletion success",
			in: storage.User{
				DeletedBy: sql.NullString{
					String: uuid.NewString(),
					Valid:  true,
				},
			},
			wantErr:  false,
			teardown: deleteTestUser,
		}, {
			name: "user deletion failed",
			in: storage.User{
				DeletedBy: sql.NullString{
					String: uuid.NewString(),
					Valid:  true,
				},
			},
			wantErr: true,
		},
	}

	uID, err := s.CreateUser(context.TODO(), storage.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password",
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			tc.in.ID = uID
			err := s.DeleteUser(context.TODO(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.DeleteUser() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr {
				t.Errorf("Storage.DeleteUser() want %v,  got %v", uID)
			}

			if tc.teardown != nil {
				tc.teardown(t, s, uID)
			}
		})
	}
}
