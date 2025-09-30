package exception

import "errors"

var (
	ErrMismatchedpassword      = errors.New("password did not match")
	ErrBodyParser              = errors.New("invalid request body")
	ErrValidation              = errors.New("validation error")
	ErrUserNotFound            = errors.New("user not found")
	ErrTokenMissingOrMalformed = errors.New("token is missing or malformed")
	ErrTokenExpired            = errors.New("token is expired")
	ErrClaimToken              = errors.New("token cannot be claim")
	ErrInvalidSubjectFormat    = errors.New("invalid subject format")
)
