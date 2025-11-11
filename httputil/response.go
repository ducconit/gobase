package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestIDHeaderKey is the default header key for the request ID.
var RequestIDHeaderKey = "X-Request-ID"

// Standard application error codes.
var (
	ErrNone               = "0"
	ErrNotFound           = "404"
	ErrUnauthorized       = "401"
	ErrForbidden          = "403"
	ErrBadRequest         = "400"
	ErrValidation         = "422"
	ErrInternalServer     = "500"
	ErrServiceUnavailable = "503"
)

// JsonResponse defines the standard structure for an API response.
type JsonResponse struct {
	// Code indicates the status of the request. "0" for success, non-zero strings for errors.
	Code string `json:"code"`

	// Message provides a human-readable message, suitable for displaying in a toast or notification.
	Message string `json:"message,omitempty"`

	// RequestID is an optional unique identifier for the request, useful for tracing and logging.
	RequestID string `json:"request_id,omitempty"`

	// Data contains the main payload of the response. It can be any type.
	Data any `json:"data,omitempty"`

	// Extra holds additional information, such as pagination details or validation errors.
	Extra any `json:"extra,omitempty"`
}

// Success sends a standardized success response (HTTP 200 OK) using Gin's context.
// It wraps the data in the standard JsonResponse struct.
func Success(c *gin.Context, data any, message string) {
	SuccessWithExtra(c, data, nil, message)
}

// SuccessWithExtra sends a standardized success response (HTTP 200 OK) with extra data using Gin's context.
func SuccessWithExtra(c *gin.Context, data any, extra any, message string) {
	response := JsonResponse{
		Code:      ErrNone,
		Message:   message,
		Data:      data,
		Extra:     extra,
		RequestID: c.Writer.Header().Get(RequestIDHeaderKey),
	}
	c.JSON(http.StatusOK, response)
}

// Error sends a standardized error response using Gin's context.
// The HTTP status code should be an appropriate error code (e.g., 400, 404, 500).
// It accepts an optional `extra` parameter for additional details.
func Error(c *gin.Context, httpStatusCode int, errorCode string, errorMessage string, extra ...any) {
	response := JsonResponse{
		Code:      errorCode,
		Message:   errorMessage,
		RequestID: c.Writer.Header().Get(RequestIDHeaderKey),
	}
	if len(extra) > 0 {
		response.Extra = extra[0]
	}
	c.JSON(httpStatusCode, response)
}

// ValidationError sends a standardized validation error response (HTTP 422) by calling the Error function.
func ValidationError(c *gin.Context, validationErrors map[string]any, message string) {
	Error(c, http.StatusUnprocessableEntity, ErrValidation, message, validationErrors)
}

// NotFound sends a 404 Not Found response.
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, ErrNotFound, message)
}

// Forbidden sends a 403 Forbidden response.
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, ErrForbidden, message)
}

// BadRequest sends a 400 Bad Request response.
func BadRequest(c *gin.Context, message string, extra ...any) {
	Error(c, http.StatusBadRequest, ErrBadRequest, message, extra...)
}

// Unauthorized sends a 401 Unauthorized response.
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, ErrUnauthorized, message)
}

// InternalServerError sends a 500 Internal Server Error response.
func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, ErrInternalServer, message)
}

// ServiceUnavailable sends a 503 Service Unavailable response to indicate service maintenance or unavailability.
func ServiceUnavailable(c *gin.Context, message string) {
	Error(c, http.StatusServiceUnavailable, ErrServiceUnavailable, message)
}
