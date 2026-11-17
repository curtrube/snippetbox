package models

import "errors"

var (
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCredentials if a user tries to log in with
	// an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrDuplicateEmail if a user tries to sign up with an email
	// that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)
