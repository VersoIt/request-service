package model

import "errors"

var (
	ErrRequestUnderConsideration = errors.New("request under consideration")
	ErrRequestNotFound           = errors.New("request not found")
)
