package errors

import "errors"

var (
	ErrOrderNotFound      = errors.New("repository not found")
	ErrInternalError      = errors.New("internal error")
	ErrSomeDetailsMissing = errors.New("some details missing")
	ErrDataBaseError      = errors.New("data base error")
)
