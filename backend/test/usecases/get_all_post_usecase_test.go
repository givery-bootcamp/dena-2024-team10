package usecases

import (
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/usecases"
	mock_interfaces "myapp/test/mock"
	"reflect"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

// リポジトリのモック
func TestGetAllPosts(t *testing.T) {

	// create test cases
	testcases := []struct {
		testName      string
		prepareMockFn func(m *mock_interfaces.MockPostRepository)
		want          []*entities.Post
		wantsFail     bool
	}{
		{
			"Success get all posts",
			func(m *mock_interfaces.MockPostRepository) {
				m.EXPECT().GetAll().Return(
					[]*entities.Post{{Id: 1, Title: "Mock", Body: "Mockやで", UserId: 2, UserName: "Mocker", CreatedAt: "2024-05-28 13:52:55", UpdatedAt: "2024-05-28 13:52:55"}},
					nil,
				)
			},
			[]*entities.Post{{Id: 1, Title: "Mock", Body: "Mockやで", UserId: 2, UserName: "Mocker", CreatedAt: "2024-05-28 13:52:55", UpdatedAt: "2024-05-28 13:52:55"}},
			false,
		},
		{
			"Fail get all posts",
			func(m *mock_interfaces.MockPostRepository) {
				m.EXPECT().GetAll().Return(
					nil,
					fmt.Errorf("Fail to get all post"),
				)
			},
			[]*entities.Post{},
			true},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockPostRepository := mock_interfaces.NewMockPostRepository(mockCtrl)
			tc.prepareMockFn(mockPostRepository)

			usecase := usecases.NewGetAllPostsUsecase(mockPostRepository)

			result, err := usecase.Execute()
			if !tc.wantsFail {
				if err != nil {
					t.Errorf("Usecase returns error: %v", err.Error())
				}
				if !reflect.DeepEqual(result, tc.want) {
					t.Errorf("Usecase returns unexpected values: %v", result)
				}
			} else {
				if err == nil {
					t.Errorf("Usecase doesn't resturn error")
				}
				if result != nil {
					t.Errorf("Usecase returns unexpected values: %v", result)
				}
			}
		})
	}
}
