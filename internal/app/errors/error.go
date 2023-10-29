package errors

import "errors"

var (
	// ErrorOccur - .
	ErrorOccur = errors.New("error")

	// ErrorNoAuthorization - .
	ErrorNoAuthorization = errors.New("authorization is not valid")

	// ErrorAuth - .
	ErrorAuth = errors.New("username or password invalid")
)
