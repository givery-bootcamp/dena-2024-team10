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
			} else {
				// test for Fail cases は作成しない
			}
		})
	}
}
