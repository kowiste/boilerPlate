package model

import "errors"

var ErrNoDocuments error = errors.New("not found")
var ErrValidation error = errors.New("validation error")
