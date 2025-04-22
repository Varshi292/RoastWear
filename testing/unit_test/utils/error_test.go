package utils_test

import (
	"github.com/Varshi292/RoastWear/internal/utils"
	"testing"
)

func TestNewErrUserExists(t *testing.T) {
	type testCase struct {
		name          string
		username      string
		errorExpected bool
	}

	tests := []testCase{
		{
			name:          "Valid: Normal Username",
			username:      "test",
			errorExpected: true,
		},
		{
			name:          "Valid: Empty Username",
			username:      "",
			errorExpected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := utils.NewErrUserExists(tc.username)
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected instantiation of ErrUserExists to pass but got nil", tc.name)
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("Test %q: expected instantiation of ErrUserExists to fail but got %v", tc.name, err)
			}
		})
	}
}

func TestNewErrEmailExists(t *testing.T) {
	type testCase struct {
		name          string
		email         string
		errorExpected bool
	}

	tests := []testCase{
		{
			name:          "Valid: Normal Email Address",
			email:         "test@test.test",
			errorExpected: true,
		},
		{
			name:          "Valid: Empty Email Address",
			email:         "",
			errorExpected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := utils.NewErrEmailExists(tc.email)
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected instantiation of ErrEmailExist to pass but got nil", tc.name)
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("Test %q: expected instantiation of ErrEmailExist to fail but got %v", tc.name, err)
			}
		})
	}
}
