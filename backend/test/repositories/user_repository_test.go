package repositories

import (
	"myapp/internal/external"
	"myapp/internal/interfaces"
	"myapp/internal/repositories"
	"testing"
)

func setupUserRepository() (interfaces.UserRepository, func()) {
	db := external.DB.Begin()
	repo := repositories.NewUserRepository(db)
	teardown := func() {
		db.Rollback()
	}
	return repo, teardown
}

func TestGetByUsername(t *testing.T) {
	repo, teardown := setupUserRepository()
	defer teardown()

	testcases := []struct {
		testName  string
		username  string
		wantsFail bool
	}{
		{"Success", "taro", false},
		// DB から意図的にエラーを返す方法がわからないため、Fail のテストケースは作成しない
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			if !tc.wantsFail {
				user, err := repo.GetByUsername(tc.username)
				if err != nil {
					t.Errorf("Repository returns error: %v", err.Error())
				}
				if user == nil {
					t.Error("Repository returns empty")
				}
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	repo, teardown := setupUserRepository()
	defer teardown()

	testcases := []struct {
		testName  string
		username  string
		password  string
		wantsFail bool
	}{
		{"Success to create user", "test_user", "test_password", false},
		{"Fail to create user", "taro", "test_password", true},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			if !tc.wantsFail {
				user, err := repo.CreateUser(tc.username, tc.password)
				if err != nil {
					t.Errorf("Repository returns error: %v", err.Error())
				}
				if user == nil {
					t.Error("Repository returns empty")
				}
			} else {
				user, err := repo.CreateUser(tc.username, tc.password)
				if err == nil {
					t.Errorf("Repository should return error")
				}
				if user != nil {
					t.Errorf(("Repository should not return user"))
				}
			}
		})
	}
}
