package service

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPostNotFound      = errors.New("post not found")
	ErrEmptyUsername     = errors.New("field can not be empty")
	ErrMessageEmpty      = errors.New("message can not be empty")
	ErrIllegalCharSymbol = errors.New("field contains illegal symbol")
)
