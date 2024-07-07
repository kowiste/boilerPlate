package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

func RestError(resp http.ResponseWriter, err error) {

	var e internalError
	ok := errors.As(err, &e)
	if !ok {
		e = internalError{
			Message: err.Error(),
			Err:     err,
			Type:    EErrorUnhandled,
			Tracing: Trace(),
		}
	}

	if e.Message == "" && e.Err != nil {
		e.Message = e.Error()
	}

	// add headers
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set(headerReason, e.Type.Message)
	resp.Header().Set(headerMessage, e.Message)
	resp.WriteHeader( e.Type.Code)

	// make and response json
	json.NewEncoder(resp).Encode(e)
}
