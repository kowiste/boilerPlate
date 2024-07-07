package errors

import (
	"net/http"
)

var (
	EErrorUnhandled = Type{
		Message:       "Unhandled Error",
		Code:         http.StatusBadRequest,
	} // 400

	EErrorValidation = Type{
		Message:       "Validation",
		Code:         http.StatusConflict,
	} // 400

	EErrorUnauthorized = Type{
		Message: "Unauthorized",
		Code:   http.StatusUnauthorized,
	} // 401

	EErrorForbidden = Type{
		Message: "Forbidden",
		Code:   http.StatusForbidden,
	} // 403

	EErrorNotFound = Type{
		Message: "Not Found",
		Code:   http.StatusNotFound,
	} // 404

	EErrorMethodNotAllowed = Type{
		Message: "Method not allowed",
		Code:   http.StatusMethodNotAllowed,
	} // 405

	EErrorNotAcceptable = Type{
		Message: "Not acceptable",
		Code:   http.StatusNotAcceptable,
	} // 406

	EErrorRequestTimeout = Type{
		Message: "Request timeout",
		Code:   http.StatusRequestTimeout,
	} // 408

	EErrorConflict = Type{
		Message: "Conflict",
		Code:   http.StatusConflict,
	} // 409

	EErrorRequestEntityTooLarge = Type{
		Message: "Request entity too large",
		Code:   http.StatusRequestEntityTooLarge,
	} // 413

	EErrorUnsupportedMediaType = Type{
		Message: "Unsupported media type",
		Code:   http.StatusUnsupportedMediaType,
	} // 415

	EErrorUpgradeRequired = Type{
		Message: "Upgrade required",
		Code:   http.StatusUpgradeRequired,
	} // 426

	EErrorUnavailableForLegalMessages = Type{
		Message: "Unavailable for legal Messages",
		Code:   http.StatusUnavailableForLegalReasons,
	} // 451

	EErrorServerInternal = Type{
		Message: "Internal Server InternalError",
		Code:   http.StatusInternalServerError,
	} // 500

	EErrorServerNotImplemented = Type{
		Message: "Not implemented",
		Code:   http.StatusNotImplemented,
	} // 501

	EErrorServerServiceUnavailable = Type{
		Message: "Service unavailable",
		Code:   http.StatusServiceUnavailable,
	} // 503

	EErrorServerHttpVersionNotSupported = Type{
		Message: "HTTP version not supported",
		Code:   http.StatusHTTPVersionNotSupported,
	} // 505
)
