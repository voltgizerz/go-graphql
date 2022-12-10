package errors

import "errors"

var (
	ErrorOccur           = errors.New("error")
	ErrorNoAuthorization = errors.New("authorization is not valid!")
	ErrorAuth            = errors.New("username or Password invalid!")
)
