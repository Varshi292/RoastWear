package bootstrap_test

import (
	"errors"
	"github.com/Varshi292/RoastWear/internal/bootstrap"
	"gorm.io/gorm"
	"testing"
)

type mockDatabase struct {
	testConnect func() (*gorm.DB, error)
	testMigrate func() error
}

func (m *mockDatabase) Connect() (*gorm.DB, error) {
	return m.testConnect()
}

func (m *mockDatabase) Migrate() error {
	return m.testMigrate()
}

func TestInitializeDatabase(t *testing.T) {
	type testCase struct {
		name        string
		connectErr  error
		migrateErr  error
		shouldPanic bool
	}

	tests := []testCase{
		{
			name:        "Successful connection and migration",
			connectErr:  nil,
			migrateErr:  nil,
			shouldPanic: false,
		},
		{
			name:        "Successful connection and failed migration",
			connectErr:  nil,
			migrateErr:  errors.New("migration error"),
			shouldPanic: true,
		},
		{
			name:        "Failed connection and successful migration",
			connectErr:  errors.New("connection error"),
			migrateErr:  nil,
			shouldPanic: true,
		},
		{
			name:        "Failed connection and failed migration",
			connectErr:  errors.New("connection error"),
			migrateErr:  errors.New("migration error"),
			shouldPanic: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := &mockDatabase{
				testConnect: func() (*gorm.DB, error) {
					if tc.connectErr != nil {
						return nil, tc.connectErr
					}
					return &gorm.DB{}, nil
				},
				testMigrate: func() error {
					return tc.migrateErr
				},
			}
			defer func() {
				if r := recover(); r != nil {
					if !tc.shouldPanic {
						t.Errorf("Unexpected panic: %v", r)
					}
				} else {
					if tc.shouldPanic {
						t.Errorf("Expected panic but none occurred")
					}
				}
			}()
			_ = bootstrap.InitializeDatabase(mockDB)
		})
	}
}
