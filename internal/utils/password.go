// Package utils contains functions for password management (hashing and verification)
// bundled to be imported wherever in the project it is necessary. The package only
// serves password managing purpose now, but is to hold other general functions
// moving forward.
package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash from a plaintext password.
//
// Parameters:
//   - password: The plaintext password to hash.
//
// Returns:
//   - string: The hashed password as a string.
//   - error: An error if hashing fails.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}
	return string(hashedPassword), nil
}

// VerifyPassword compares a bcrypt hashed password to a plaintext password.
// Returns true if the password matches, and false otherwise.
//
// Parameters:
//   - password: The plaintext password
//   - hashedPassword: The hashed password
//
// Returns:
//   - bool: True if the password matches the hash, otherwise false.
func VerifyPassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
