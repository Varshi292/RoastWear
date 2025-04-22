package utils_test

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func initializeMockDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.Session{}); err != nil {
		return nil, err
	}
	return db, nil
}

func TestCreateSession(t *testing.T) {
	db, err := initializeMockDB()
	if err != nil {
		t.Fatal("failed to create test database")
	}
	type testCase struct {
		name          string
		session       *models.Session
		createFunc    func() error
		errorExpected bool
	}

	tests := []testCase{
		{
			name: "Valid: Create Session",
			session: &models.Session{
				SessionKey: "test-test-test",
			},
			createFunc: func() error {
				repo := repositories.NewSessionRepository(db)
				return repo.CreateSession(&models.Session{SessionKey: "test-test-test"})
			},
			errorExpected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.createFunc()
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected creation to fail but it passed", tc.name)
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("Test %q: expected creation to pass but it failed: %v", tc.name, err)
			}
		})
	}
}

func TestGetSession(t *testing.T) {
	db, err := initializeMockDB()
	if err != nil {
		t.Fatal("failed to create test database")
	}
	type testCase struct {
		name          string
		sessionID     string
		getFunc       func() error
		doesExist     bool
		errorExpected bool
	}

	tests := []testCase{
		{
			name:      "Valid: Get Existing Session",
			sessionID: "test-test-test",
			getFunc: func() error {
				repo := repositories.NewSessionRepository(db)
				return repo.GetSession("test-test-test")
			},
			doesExist:     true,
			errorExpected: false,
		},
		{
			name:      "Invalid: Get Non-Existing Session",
			sessionID: "does-not-exist",
			getFunc: func() error {
				repo := repositories.NewSessionRepository(db)
				return repo.GetSession("does-not-exist")
			},
			doesExist:     false,
			errorExpected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.doesExist {
				db.Create(&models.Session{SessionKey: tc.sessionID})
			}
			err := tc.getFunc()
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected retrieval to fail but it passed", tc.name)
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("Test %q: expected retrieval to pass but it failed: %v", tc.name, err)
			}
		})
	}
}

func TestDeleteSession(t *testing.T) {
	db, err := initializeMockDB()
	if err != nil {
		t.Fatal("failed to create test database")
	}
	db, err = initializeMockDB()
	type testCase struct {
		name          string
		sessionID     string
		deleteFunc    func() error
		doesExist     bool
		errorExpected bool
	}

	tests := []testCase{
		{
			name:      "Valid: Delete Existing Session",
			sessionID: "delete-test-test",
			deleteFunc: func() error {
				repo := repositories.NewSessionRepository(db)
				return repo.DeleteSession("delete-test-test")
			},
			doesExist:     true,
			errorExpected: false,
		},
		{
			name:      "Invalid: Delete Non-Existing Session",
			sessionID: "does-not-exist",
			deleteFunc: func() error {
				repo := repositories.NewSessionRepository(db)
				return repo.DeleteSession("does-not-exist")
			},
			doesExist:     false,
			errorExpected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.doesExist {
				db.Create(&models.Session{SessionKey: tc.sessionID})
			}
			err := tc.deleteFunc()
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected deletion to fail but it passed", tc.name)
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("Test %q: expected deletion to pass but it failed: %v", tc.name, err)
			}
		})
	}
}
