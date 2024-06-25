package repositories

import (
	"myapp/internal/external"
	"myapp/internal/interfaces"
	"myapp/internal/repositories"
	"testing"
)

func setupPostRepository() (interfaces.PostRepository, func()) {
	db := external.DB.Begin()
	repo := repositories.NewPostRepository(db)
	teardown := func() {
		db.Rollback()
	}
	return repo, teardown
}

func TestGetAll(t *testing.T) {
	// initialize DB
	repo, teardown := setupPostRepository()
	defer teardown()

	// create test cases
	testcases := []struct {
		testName  string
		wantsFail bool
	}{
		// DB から意図的にエラーを返す方法がわからないため、Fail のテストケースは作成しない
		{"Success", false},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			if !tc.wantsFail {
				// test for Success cases
				posts, err := repo.GetAll()
				if err != nil {
					t.Errorf("Repository returns error: %v", err.Error())
				}
				if len(posts) == 0 {
					// DB にシードデータがない場合は、このエラーが発生する
					t.Error("Repository returns empty")
				}
			}
			// test for Fail cases は作成しない
		})
	}
}

func TestCreatePost(t *testing.T) {
	repo, teardown := setupPostRepository()
	defer teardown()

	testCases := []struct {
		name        string
		title       string
		body        string
		userId      int64
		expectError bool
	}{
		{"ValidPost", "Test Title", "Test Body", 1, false},
		{"EmptyTitle", "", "Test Body", 1, true},
		{"EmptyBody", "Test Title", "", 1, true},
		{"InvalidUserId", "Test Title", "Test Body", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.CreatePost(tc.title, tc.body, tc.userId)
			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got one: %v", err)
				}
				// Additional checks can be added here to validate the created post
			}
		})
	}
}
