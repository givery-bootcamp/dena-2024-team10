package repositories

import (
	"myapp/internal/entities"
	"myapp/internal/external"
	"myapp/internal/interfaces"
	"myapp/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupCommentRepository() (interfaces.CommentRepository, func()) {
	db := external.DB.Begin()
	repo := repositories.NewCommentRepository(db)
	teardown := func() {
		db.Rollback()
	}
	return repo, teardown
}

func TestCreate(t *testing.T) {
	// initialize DB
	repo, teardown := setupCommentRepository()
	defer teardown()

	// create test cases
	testcases := []struct {
		testName        string
		postId          int64
		userId          int64
		body            string
		expectedComment *entities.Comment
		expectedError   error
	}{
		{
			"Success",
			1,
			1,
			"test",
			&entities.Comment{
				// Id is not checked because it is set by the database
				PostId: 1,
				UserId: 1,
				Body:   "test",
				// CreatedAt and UpdatedAt are not checked because they are set by the database
			},
			nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := repo.Create(tc.postId, tc.body, tc.userId)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedComment.PostId, result.PostId)
			assert.Equal(t, tc.expectedComment.UserId, result.UserId)
			assert.Equal(t, tc.expectedComment.Body, result.Body)
		})
	}
}
