package repositories

import (
	"errors"
	"log"
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

func TestCreateComment(t *testing.T) {
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
		{
			"Fail with unexisting post",
			100,
			1,
			"test",
			nil,
			errors.New("post or user not found"),
		},
		{
			"Fail with unexisting user",
			1,
			100,
			"test",
			nil,
			errors.New("post or user not found"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := repo.Create(tc.postId, tc.body, tc.userId)
			assert.Equal(t, tc.expectedError, err)
			if result != nil {
				// Id is not checked because it is set by the database
				assert.Equal(t, tc.expectedComment.PostId, result.PostId)
				assert.Equal(t, tc.expectedComment.UserId, result.UserId)
				assert.Equal(t, tc.expectedComment.Body, result.Body)
				// CreatedAt and UpdatedAt are not checked because they are set by the database
			}
		})
	}
}

func TestGetCommentByID(t *testing.T) {
	// initialize DB
	repo, teardown := setupCommentRepository()
	defer teardown()

	// create test cases
	testcases := []struct {
		testName        string
		commentId       int64
		expectedComment *entities.Comment
		expectedError   error
	}{
		{
			"Success",
			1,
			&entities.Comment{ // defined in the seed
				Id:       1,
				PostId:   1,
				UserId:   1,
				Body:     "comment1 on test1",
				Username: "taro",
				// CreatedAt and UpdatedAt are not checked
				// because they are set by the database.
			},
			nil,
		},
		{
			"Success with unexisting comment",
			100,
			nil,
			nil,
		},
		// Do not test the case where DB returns an unknown error
		// because it is difficult to reproduce.
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := repo.GetById(tc.commentId)
			assert.Equal(t, tc.expectedError, err)
			if result != nil {
				assert.Equal(t, tc.expectedComment.Id, result.Id)
				assert.Equal(t, tc.expectedComment.PostId, result.PostId)
				assert.Equal(t, tc.expectedComment.UserId, result.UserId)
				assert.Equal(t, tc.expectedComment.Body, result.Body)
				assert.Equal(t, tc.expectedComment.Username, result.Username)
				// do not check CreatedAt and UpdatedAt because they are set by the database
			}
		})
	}
}

func TestUpdateComment(t *testing.T) {
	// initialize DB
	repo, teardown := setupCommentRepository()
	defer teardown()

	// create test cases
	testcases := []struct {
		testName        string
		inputComment    *entities.Comment
		body            string
		expectedComment *entities.Comment
		expectedError   error
	}{
		{
			"Success",
			&entities.Comment{ // defined in the seed
				Id:     1,
				PostId: 1,
				UserId: 1,
				Body:   "updated comment1 on test1",
				// CreatedAt and UpdatedAt are not checked
				// because they are set by the database.
			},
			"updated comment1 on test1",
			&entities.Comment{
				Id:     1,
				PostId: 1,
				UserId: 1,
				Body:   "updated comment1 on test1",
				// CreatedAt and UpdatedAt are not checked
				// because they are set by the database.
			},
			nil,
		},
		{
			"Fail with unexisting comment",
			&entities.Comment{
				Id:     100,
				PostId: 1,
				UserId: 1,
				Body:   "test",
			},
			"test",
			nil,
			errors.New("comment not found"),
		},
		// Do not test the case where DB returns an unknown error
		// because it is difficult to reproduce.
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := repo.Update(tc.inputComment)
			assert.Equal(t, tc.expectedError, err)
			if result != nil {
				assert.Equal(t, tc.expectedComment.Id, result.Id)
				assert.Equal(t, tc.expectedComment.PostId, result.PostId)
				assert.Equal(t, tc.expectedComment.UserId, result.UserId)
				assert.Equal(t, tc.expectedComment.Body, result.Body)
				// do not check CreatedAt and UpdatedAt because they are set by the database
			}
		})
	}
}

func TestDeleteComment(t *testing.T) {
	// initialize DB
	repo, teardown := setupCommentRepository()
	defer teardown()

	// create test cases
	testcases := []struct {
		testName      string
		commentId     int64
		expectedError error
	}{
		{
			"Success",
			1,
			nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			err := repo.Delete(tc.commentId)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetCommentByPostId(t *testing.T) {
	// initialize DB
	repo, teardown := setupCommentRepository()
	defer teardown()

	// create test cases
	testcases := []struct {
		testName         string
		postId           int64
		limit            int64
		offset           int64
		expectedComments []*entities.Comment
		expectedError    error
	}{
		{
			"Success limit=1 offset=0",
			1,
			1,
			0,
			[]*entities.Comment{
				{ // defined in the seed
					Id:       1,
					PostId:   1,
					UserId:   1,
					Body:     "comment1 on test1",
					Username: "taro",
					// CreatedAt and UpdatedAt are not checked
					// because they are set by the database.
				},
			},
			nil,
			// Do not test the case where DB returns an unknown error
			// because it is difficult to reproduce.
		},
		{
			"Success limit=2 offset=1",
			1,
			2,
			1,
			[]*entities.Comment{
				{ // defined in the seed
					Id:       2,
					PostId:   1,
					UserId:   2,
					Body:     "comment2 on test1",
					Username: "hanako",
				},
				{ // defined in the seed
					Id:       3,
					PostId:   1,
					UserId:   1,
					Body:     "comment3 on test1",
					Username: "taro",
				},
			},
			nil,
			// Do not test the case where DB returns an unknown error
			// because it is difficult to reproduce.
		},
		{
			"Success limit=1 offset=25 (return empty)",
			1,
			1,
			25,
			[]*entities.Comment{},
			nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := repo.GetAllByPostId(tc.postId, tc.limit, tc.offset)
			assert.Equal(t, tc.expectedError, err)

			for i, comment := range result {
				log.Println(comment)
				assert.Equal(t, tc.expectedComments[i].Id, comment.Id)
				assert.Equal(t, tc.expectedComments[i].PostId, comment.PostId)
				assert.Equal(t, tc.expectedComments[i].UserId, comment.UserId)
				assert.Equal(t, tc.expectedComments[i].Body, comment.Body)
				assert.Equal(t, tc.expectedComments[i].Username, comment.Username)

			}
		})
	}
}
