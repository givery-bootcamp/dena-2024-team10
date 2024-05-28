package repositories

import (
	"myapp/internal/external"
	"myapp/internal/interfaces"
	"myapp/internal/repositories"
	"testing"
)

func init() {
	external.SetupDB()
}

func setupHelloWorld() (interfaces.HelloWorldRepository, func()) {
	db := external.DB.Begin()
	repo := repositories.NewHelloWorldRepository(db)
	teardown := func() {
		db.Rollback()
	}
	return repo, teardown
}

func TestHelloWorld(t *testing.T) {
	// initialize DB
	repo, teardown := setupHelloWorld()
	defer teardown()

	// create test cases
	testcases := []struct {
		testName  string
		lang      string
		message   string
		wantsFail bool
	}{
		{"Success with English", "en", "Hello World", false},
		{"Success with Japanese", "ja", "こんにちは 世界", false},
		{"Fail with French", "fr", "", true},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			if !tc.wantsFail {
				// test for Success cases
				result, err := repo.Get(tc.lang)
				if err != nil {
					t.Errorf("Repository returns error: %v", err.Error())
				}
				if result == nil {
					t.Error("Nil")
				} else if result.Message != tc.message {
					t.Errorf("Wrong value: %+v", result)
				}
			} else {
				// test for Fail cases
				result, err := repo.Get(tc.lang)
				if err != nil {
					t.Errorf("Repository returns error: %v", err.Error())
				}

				if result != nil {
					t.Errorf("Not nil %+v", result)
				}
			}
		})
	}
}
