package shared_errors

import "errors"

var (
	ErrCannotOpenGraphQLFile = errors.New("unable to open graphql schema file")
	ErrCannotReadGraphQLFile = errors.New("failed to read graphql schema file")
	ErrUnableToMarshalToJson = errors.New("unable to marshal from JSON")
	ErrRequestedUserNotFound = errors.New("requested user not found")
)
