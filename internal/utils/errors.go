package utils

import (
	"errors"
	"fmt"
)

type ErrUserExists struct {
	Username string
}

func (e *ErrUserExists) Error() string {
	return fmt.Sprintf("user '%s' already exists", e.Username)
}
func NewErrUserExists(username string) error {
	return &ErrUserExists{Username: username}
}

type ErrEmailExists struct {
	Email string
}

func (e *ErrEmailExists) Error() string {
	return fmt.Sprintf("email address '%s' already exists", e.Email)
}
func NewErrEmailExists(email string) error {
	return &ErrEmailExists{Email: email}
}

var ErrInvalidCredentials = errors.New("invalid username or password")
var ErrSessionNotFound = errors.New("session not found")
