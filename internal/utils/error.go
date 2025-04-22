package utils

import (
	"errors"
	"fmt"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type ErrUserExists struct {
	Username string
}

func NewErrUserExists(username string) error {
	return &ErrUserExists{Username: username}
}
func (e *ErrUserExists) Error() string {
	return fmt.Sprintf("user '%s' already exists", e.Username)
}

type ErrEmailExists struct {
	Email string
}

func NewErrEmailExists(email string) error {
	return &ErrEmailExists{Email: email}
}
func (e *ErrEmailExists) Error() string {
	return fmt.Sprintf("email address '%s' already exists", e.Email)
}
