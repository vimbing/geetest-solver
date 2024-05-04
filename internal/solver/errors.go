package solver

import "errors"

var (
	ErrResponseInvalid    = errors.New("response is not valid")
	ErrImageNotRecognized = errors.New("image is not recognized")
)
