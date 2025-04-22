package repositories_test

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
	if err := db.Create(&models.Session{SessionKey: "test-test-test"}).Error; err != nil {
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
		input         *models.Session
		errorExpected bool
	}
	repo := repositories.NewSessionRepository(db)

	tests := []testCase{
		{
			name:          "Valid: New Session Inserted",
			input:         &models.Session{SessionKey: "new-session-1"},
			errorExpected: false,
		},
		{
			name:          "Invalid: Duplicate SessionKey",
			input:         &models.Session{SessionKey: "test-test-test"},
			errorExpected: true,
		},
		{
			name:          "Valid: Empty SessionKey",
			input:         &models.Session{SessionKey: ""},
			errorExpected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.CreateSession(tc.input)
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected creation to fail but it passed (input: %+v)", tc.name, tc.input)
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("Test %q: expected creation to pass but it failed: %v (input: %+v)", tc.name, err, tc.input)
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
		input         string
		errorExpected bool
	}

	repo := repositories.NewSessionRepository(db)

	tests := []testCase{
		{
			name:          "Valid: Session Exists",
			input:         "test-test-test",
			errorExpected: false,
		},
		{
			name:          "Invalid: Nonexistent Session",
			input:         "non-existent-session",
			errorExpected: true,
		},
		{
			name:          "Invalid: Empty Session ID",
			input:         "",
			errorExpected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.GetSession(tc.input)
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected get to fail but it passed (input: %s)", tc.name, tc.input)
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("Test %q: expected get to pass but it failed: %v (input: %s)", tc.name, err, tc.input)
			}
		})
	}
}

func TestDeleteSession(t *testing.T) {
	db, err := initializeMockDB()
	if err != nil {
		t.Fatal("failed to create test database")
	}

	type testCase struct {
		name          string
		input         string
		errorExpected bool
	}

	repo := repositories.NewSessionRepository(db)

	tests := []testCase{
		{
			name:          "Valid: Delete Existing Session",
			input:         "test-test-test",
			errorExpected: false,
		},
		{
			name:          "Invalid: Delete Nonexistent Session",
			input:         "non-existent-session",
			errorExpected: false,
		},
		{
			name:          "Invalid: Empty Session ID",
			input:         "",
			errorExpected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.DeleteSession(tc.input)
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected deletion to fail but it passed (input: %s)", tc.name, tc.input)
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("Test %q: expected deletion to pass but it failed: %v (input: %s)", tc.name, err, tc.input)
			}
		})
	}
}
