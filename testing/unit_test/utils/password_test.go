package utils_test

import (
	"github.com/Varshi292/RoastWear/internal/utils"
	"testing"
)

func TestHashPassword(t *testing.T) {
	type testCase struct {
		name          string
		input         string
		errorExpected bool
	}

	tests := []testCase{
		{
			name:          "Valid: Normal Password",
			input:         "password",
			errorExpected: false,
		},
		{
			name:          "Valid: Empty Password",
			input:         "",
			errorExpected: false,
		},
		{
			name:          "Valid: Password With Numbers",
			input:         "pa55w0rd",
			errorExpected: false,
		},
		{
			name:          "Valid: Password With Special Characters",
			input:         "p@ssw@rd",
			errorExpected: false,
		},
		{
			name:          "Valid: Password With Whitespace",
			input:         " password ",
			errorExpected: false,
		},
		{
			name:          "Valid: Password With Only Whitespace",
			input:         "                      ",
			errorExpected: false,
		},
		{
			name:          "Valid: Password With Unicode",
			input:         "passw🌍rd",
			errorExpected: false,
		},
		{
			name:          "Valid: Maximum-length (72 bytes) Password",
			input:         "passssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssword",
			errorExpected: false,
		},
		{
			name:          "Invalid: Maximum Password Length Exceeded (73 bytes)",
			input:         "passsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssword",
			errorExpected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hashed, err := utils.HashPassword(tc.input)
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected hashing to fail but it passed (input: %q)", tc.name, tc.input)
			}
			if !tc.errorExpected {
				if err != nil {
					t.Errorf("Test %q: expected hashing to pass but it failed: %v (input: %q)", tc.name, err, tc.input)
				}
				if hashed == "" {
					t.Errorf("Test %q: expected non-empty hash result but received empty string", tc.name)
				}
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type testCase struct {
		name          string
		password      string
		hashFunc      func(string) string
		errorExpected bool
	}

	validPassword := "correctPassword?"
	validHash, err := utils.HashPassword(validPassword)
	if err != nil {
		t.Fatalf("Failed to generate valid hash for setup: %v", err)
	}

	tests := []testCase{
		{
			name:          "Valid: Correct Password and Hash",
			password:      validPassword,
			hashFunc:      func(h string) string { return h },
			errorExpected: false,
		},
		{
			name:          "Invalid: Incorrect Password",
			password:      "incorrectPassword!",
			hashFunc:      func(h string) string { return h },
			errorExpected: true,
		},
		{
			name:          "Invalid: Empty Password",
			password:      "",
			hashFunc:      func(h string) string { return h },
			errorExpected: true,
		},
		{
			name:          "Invalid: Incorrect Hash",
			password:      validPassword,
			hashFunc:      func(_ string) string { return "correctPassword@TrustMe" },
			errorExpected: true,
		},
		{
			name:          "Invalid: Empty Hash",
			password:      validPassword,
			hashFunc:      func(_ string) string { return "" },
			errorExpected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.VerifyPassword(tc.password, tc.hashFunc(validHash))
			if tc.errorExpected && result {
				t.Errorf("Test %q: expected verification to fail but it passed (password: %q)", tc.name, tc.password)
			}
			if !tc.errorExpected && !result {
				t.Errorf("Test %q: expected verification to pass but it failed (password: %q)", tc.name, tc.password)
			}
		})
	}
}
