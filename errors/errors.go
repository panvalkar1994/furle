package errors

import "errors"

var (
	ErrUrlRequired = errors.New("URL is required")
)

type ErrorPage struct {
	Error    error  `json:"error"`
	ErrorMsg string `json:"error_msg"`
}
