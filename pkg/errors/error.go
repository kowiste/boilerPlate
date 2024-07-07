package errors

import "errors"

type internalError struct {
	Message string         `json:"message"`
	Err     error          `json:"-"`
	Type    Type           `json:"type,omitempty"`
	Tracing trace          `json:"trace,omitempty"`
	Data    map[string]any `json:"-"`
}

func (i internalError) Error() string {
	if i.Err == nil {
		return i.Message
	}

	return i.Err.Error()
}

func New(
	message string,
	Type Type,
	data ...map[string]any,
) internalError {
	ie := internalError{
		Message: message,
		Err:     errors.New(message),
		Tracing: Trace(),
		Type: Type,
	}

	if len(data) > 0 {
		ie.Data = data[0]
	}

	return ie
}