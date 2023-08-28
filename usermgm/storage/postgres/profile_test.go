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

func insertTestUserInformation(t *testing.T, s *Storage) {
	// insert test user into the "users" table
	got, err := s.CreateUserInformation(context.TODO(), storage.UserInformation{
		UserID:    "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
	})
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}

	if got == "" {
		t.Fatalf("Unable to create user: %v", err)
	}
}

func deleteTestUserInformation(t *testing.T, s *Storage, id string) {
	// delete test user from the "users" table
	if err := s.DeleteUserInformation(context.TODO(), id, id); err != nil {
		t.Fatalf("Unable to delete user: %v", err)
	}
}

func deleteUserInformationPermanently(t *testing.T, s *Storage) {
	// delete test user from the "users" table
	if err := s.DeleteUserInformationPermanently(context.TODO()); err != nil {
		t.Fatalf("Unable to delete users: %v", err)
	}
}

func TestCreateUserInformation(t *testing.T) {
	s := newTestStorage(t)
	testCases := []struct {
		name     string
		in       storage.UserInformation
		want     string
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage)
	}{
		{
			name: "user information creation success",
			in: storage.UserInformation{
				Image:     "1.jpg",
				FirstName: "Super",
				LastName:  "Admin",
				Mobile:    "01715039303",
				Gender:    1,
				Address:   "Farazipara",
				City:      "Khulna",
				Country:   "Bangladesh",
				UserID:    "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
			},
			wantErr:  false,
			teardown: deleteUserInformationPermanently,
		}, {
			name: "user information creation failed",
			in: storage.UserInformation{
				Image:     "017",
				FirstName: "",
				LastName:  "Admin",
				Mobile:    "01715039303",
				Gender:    1,
				Address:   "Farazipara",
				City:      "Khulna",
				Country:   "Bangladesh",
				UserID:    "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
			},
			wantErr:  true,
			setup:    insertTestUserInformation,
			teardown: deleteUserInformationPermanently,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			got, err := s.CreateUserInformation(context.TODO(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.CreateUserInformation() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && got == "" {
				t.Errorf("Storage.CreateUserInformation() want uuid, got %v", got)
				return
			}

			if tc.teardown != nil {
				tc.teardown(t, s)
			}
		})
	}
}

var UserInformation = []struct {
	name     string
	id       string
	want     *storage.UserInformation
	wantErr  bool
	teardown func(*testing.T, *Storage, string)
	setup    func(*testing.T, *Storage, string)
}{
	{
		name: "Get User Information Successfully",
		id:   "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
		want: &storage.UserInformation{
			Image:     "default.jpg",
			FirstName: "Super",
			LastName:  "Admin",
			Mobile:    "+8801715039303",
			Gender:    1,
			Address:   "farazipara",
			City:      "khulna",
			Country:   "bangladesh",
			UserID:    "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
		},
		teardown: deleteTestUser,
	}, {
		name:    "Get User Information Failed",
		id:      "b6ddbe32-3d7e-4828-b2d7-da9927ue8ry",
		wantErr: true,
	},
}

func TestGetUserInformation(t *testing.T) {
	ts := newTestStorage(t)
	_, err := ts.db.Exec("INSERT INTO user_information (image, first_name, last_name, mobile, address, city, country, user_id) VALUES ('default.jpg', 'Super', 'Admin', '+8801715039303', 'farazipara', 'khulna', 'bangladesh', 'b6ddbe32-3d7e-4828-b2d7-da9927846e6b')")
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.UserInformation{},
			"ID",
			"DOB",
			"CreatedAt",
			"CreatedBy",
			"UpdatedAt",
			"UpdatedBy",
			"DeletedAt",
			"DeletedBy",
		),
	}

	for _, tc := range UserInformation {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ts.GetUserInformation(context.Background(), tc.id)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.Profile() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.Profile() diff = %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, ts, tc.id)
			}
		})
	}
}

var testUpdateProfile = []struct {
	name     string
	in       storage.UserInformation
	want     *storage.UserInformation
	wantErr  bool
	setup    func(*testing.T, *Storage)
	teardown func(*testing.T, *Storage, string)
}{
	{
		name: "update profile success",
		in: storage.UserInformation{
			Image:     "2.jpg",
			FirstName: "Super",
			LastName:  "Admin",
			Mobile:    "01715039303",
			Gender:    2,
			Address:   "Farazipara",
			City:      "Khulna",
			Country:   "Bangladesh",
			UserID:    "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
		},
		want: &storage.UserInformation{
			Image:     "2.jpg",
			FirstName: "Super",
			LastName:  "Admin",
			Mobile:    "01715039303",
			Gender:    2,
			Address:   "Farazipara",
			City:      "Khulna",
			Country:   "Bangladesh",
			UserID:    "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
		},
		teardown: deleteTestUserInformation,
	},
}

func TestUpdateProfile(t *testing.T) {
	ts := newTestStorage(t)
	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.UserInformation{},
			"ID",
			"DOB",
			"CreatedAt",
			"CreatedBy",
			"UpdatedAt",
			"UpdatedBy",
			"DeletedAt",
			"DeletedBy",
		),
	}

	for _, tc := range testUpdateProfile {
		t.Run(tc.name, func(t *testing.T) {
			tc.in.UpdatedBy = sql.NullString{String: tc.in.UserID, Valid: true}
			got, err := ts.UpdateUserInformation(context.Background(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.UpdateProfile() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.UpdateProfile() diff = %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, ts, tc.in.UserID)
			}
		})
	}
}

func TestDeleteUserInformation(t *testing.T) {
	s := newTestStorage(t)
	testCases := []struct {
		name     string
		in       storage.UserInformation
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage, string)
	}{
		{
			name: "user information deletion success",
			in: storage.UserInformation{
				UserID: "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
				DeletedBy: sql.NullString{
					String: uuid.NewString(),
					Valid:  true,
				},
			},
			wantErr:  false,
			teardown: deleteTestUserInformation,
		}, {
			name: "user information deletion failed",
			in: storage.UserInformation{
				DeletedBy: sql.NullString{
					String: uuid.NewString(),
					Valid:  true,
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			err := s.DeleteUserInformation(context.TODO(), tc.in.UserID, tc.in.UserID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.DeleteUserInformation() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
