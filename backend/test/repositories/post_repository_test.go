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
		testName      string
		limit         int64
		offset        int64
		expectedPosts []*entities.Post
		expectedError error
	}{
		{
			"Success limit=1 offset=0",
			1,
			0,
			[]*entities.Post{
				{
					Id:       1,
					Title:    "test1",
					Body:     "質問1\n改行",
					UserId:   1,
					UserName: "taro",
				},
			},
			nil,
		},
		{
			"Success limit=2 offset=1",
			2,
			1,
			[]*entities.Post{
				{
					Id:       2,
					Title:    "test2",
					Body:     "質問2\n改行",
					UserId:   1,
					UserName: "taro",
				},
				{
					Id:       3,
					Title:    "test3",
					Body:     "質問3\n改行",
					UserId:   2,
					UserName: "hanako",
				},
			},
			nil,
		},
		// DB から意図的にエラーを返す方法がわからないため、Fail のテストケースは作成しない
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			posts, err := repo.GetAll(tc.limit, tc.offset)
			assert.Equal(t, tc.expectedError, err)
			assert.Len(t, posts, len(tc.expectedPosts))
			for i, post := range posts {
				assert.Equal(t, tc.expectedPosts[i].Id, post.Id)
				assert.Equal(t, tc.expectedPosts[i].Title, post.Title)
				assert.Equal(t, tc.expectedPosts[i].Body, post.Body)
				assert.Equal(t, tc.expectedPosts[i].UserId, post.UserId)
				assert.Equal(t, tc.expectedPosts[i].UserName, post.UserName)
			}
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
				Id:       1,
				Title:    "test1",
				Body:     "質問1\n改行",
				UserId:   1,
				UserName: "taro",
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
				assert.Equal(t, tc.expectedPost.UserName, post.UserName)
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
