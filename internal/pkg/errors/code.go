// internal/pkg/errors/code.go
package errors

import "net/http"

type ErrorCode int

const (
	OK ErrorCode = iota
	InvalidParameter
	Unauthorized
	Forbidden
	NotFound
	InternalError
	RateLimitExceeded
)

var errorMessages = map[ErrorCode]string{
	OK:                "success",
	InvalidParameter:  "invalid parameter",
	Unauthorized:      "unauthorized",
	Forbidden:         "forbidden",
	NotFound:          "not found",
	InternalError:     "internal server error",
	RateLimitExceeded: "rate limit exceeded",
}

func (e ErrorCode) Message() string {
	return errorMessages[e]
}

func (e ErrorCode) Status() int {
	switch e {
	case OK:
		return http.StatusOK
	case InvalidParameter:
		return http.StatusBadRequest
	case Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	case NotFound:
		return http.StatusNotFound
	case InternalError:
		return http.StatusInternalServerError
	case RateLimitExceeded:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}
