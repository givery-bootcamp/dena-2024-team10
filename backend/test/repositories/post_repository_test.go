package repositories

import (
	"myapp/internal/entities"
	"myapp/internal/external"
	"myapp/internal/interfaces"
	"myapp/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestGetById(t *testing.T) {
	// initialize DB
	repo, teardown := setupPostRepository()
	defer teardown()

	// create test cases
	testcases := []struct {
		testName     string
		postId       int64
		expectedPost *entities.Post
		expectedErr  error
	}{
		{
			"Success",
			1,
			&entities.Post{ // migration で作成されたデータ
				Id:     1,
				Title:  "test1",
				Body:   "質問1\n改行",
				UserId: 1,
				// CreatedAt などはテストケースに含めない
			},
			nil,
		},
		{
			"Not Found",
			100,
			nil,
			nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			post, err := repo.GetById(tc.postId)
			assert.Equal(t, tc.expectedErr, err)
			// post が nil の場合は、assert.Equal でエラーが発生するため、if 文で分岐
			if post != nil {
				assert.Equal(t, tc.expectedPost.Id, post.Id)
				assert.Equal(t, tc.expectedPost.Title, post.Title)
				assert.Equal(t, tc.expectedPost.Body, post.Body)
				assert.Equal(t, tc.expectedPost.UserId, post.UserId)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	// initialize DB
	repo, teardown := setupPostRepository()
	defer teardown()

	// create test cases
	testcases := []struct {
		testName  string
		postId    int64
		wantsFail bool
	}{
		{
			"Success",
			1,
			false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			err := repo.Delete(tc.postId)
			if tc.wantsFail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
