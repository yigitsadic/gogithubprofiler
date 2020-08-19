package shared

import "errors"

var (
	ErrUnableToMarshalToJson = errors.New("unable to marshal from JSON")
	ErrRequestedUserNotFound = errors.New("requested user not found")
)
