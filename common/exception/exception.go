package exception

import "errors"

var (
	ErrMismatchedpassword = errors.New("password did not match")
	ErrBodyParser         = errors.New("invalid request body")
	ErrValidation         = errors.New("validation error")
	ErrUserNotFound       = errors.New("user not found")
)
