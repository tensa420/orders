package entity

import "errors"

var (
	ErrOrderNotFound      = errors.New("repository not found")
	ErrInternalError      = errors.New("internal error")
	ErrSuccessCancel      = errors.New("success cancel")
	ErrSomeDetailsMissing = errors.New("some details missing")
)
