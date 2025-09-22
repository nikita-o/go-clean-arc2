package exception

import (
	"errors"
)

type BackendErrorCode string

const (
	UnknownErrorCode  BackendErrorCode = "Unknown"
	NotFoundErrorCode BackendErrorCode = "NotFound"
)

type BackendError error

var (
	UnknownError  BackendError = errors.New("unknown error")
	NotFoundError BackendError = errors.New("not found")
)
