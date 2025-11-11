package httputil

import (
	"net/http"

	"github.com/ducconit/gobase/paginate"
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

// JsonResponse defines the standard structure for an API response with generic types.
type JsonResponse[T any, E any] struct {
	// Code indicates the status of the request. "0" for success, non-zero strings for errors.
	Code string `json:"code"`

	// Message provides a human-readable message, suitable for displaying in a toast or notification.
	Message string `json:"message,omitempty"`

	// RequestID is an optional unique identifier for the request, useful for tracing and logging.
	RequestID string `json:"request_id,omitempty"`

	// Data contains the main payload of the response.
	Data T `json:"data,omitempty"`

	// Extra holds additional information, such as pagination details or validation errors.
	Extra E `json:"extra,omitempty"`
}

// Success sends HTTP 200 with data and message.
func Success[T any](c *gin.Context, data T, message string) {
	SuccessWithExtra[T, any](c, data, nil, message)
}

// SuccessWithExtra sends HTTP 200 with data, extra metadata, and message.
func SuccessWithExtra[T any, E any](c *gin.Context, data T, extra E, message string) {
	resp := JsonResponse[T, E]{
		Code:      ErrNone,
		Message:   message,
		Data:      data,
		Extra:     extra,
		RequestID: c.Writer.Header().Get(RequestIDHeaderKey),
	}
	c.JSON(http.StatusOK, resp)
}

// Error sends an error response with HTTP status code and optional extra data.
func Error[E any](c *gin.Context, httpStatusCode int, errorCode string, errorMessage string, extra ...E) {
	var extraData E
	if len(extra) > 0 {
		extraData = extra[0]
	}
	resp := JsonResponse[any, E]{
		Code:      errorCode,
		Message:   errorMessage,
		RequestID: c.Writer.Header().Get(RequestIDHeaderKey),
		Extra:     extraData,
	}
	c.JSON(httpStatusCode, resp)
}

// ValidationError sends a 422 Unprocessable Entity response with validation errors.
func ValidationError[E map[string]any](c *gin.Context, validationErrors E, message string) {
	Error(c, http.StatusUnprocessableEntity, ErrValidation, message, validationErrors)
}

// NotFound sends a 404 Not Found response.
func NotFound(c *gin.Context, message string) {
	Error[any](c, http.StatusNotFound, ErrNotFound, message)
}

// Forbidden sends a 403 Forbidden response.
func Forbidden(c *gin.Context, message string) {
	Error[any](c, http.StatusForbidden, ErrForbidden, message)
}

// BadRequest sends a 400 Bad Request response with optional extra data.
func BadRequest[E any](c *gin.Context, message string, extra ...E) {
	Error(c, http.StatusBadRequest, ErrBadRequest, message, extra...)
}

// Unauthorized sends a 401 Unauthorized response.
func Unauthorized(c *gin.Context, message string) {
	Error[any](c, http.StatusUnauthorized, ErrUnauthorized, message)
}

// InternalServerError sends a 500 Internal Server Error response.
func InternalServerError(c *gin.Context, message string) {
	Error[any](c, http.StatusInternalServerError, ErrInternalServer, message)
}

// ServiceUnavailable sends a 503 Service Unavailable response.
func ServiceUnavailable(c *gin.Context, message string) {
	Error[any](c, http.StatusServiceUnavailable, ErrServiceUnavailable, message)
}

// SimplePagination sends HTTP 200 with items and simple pagination metadata.
func SimplePagination[T any](c *gin.Context, items []T, total int64, page int, pageSize int, message string) {
	sp := paginate.NewSimplePagination(total, page, pageSize)
	SuccessWithExtra[[]T, *paginate.SimplePagination](c, items, sp, message)
}

// CursorPagination sends HTTP 200 with items and cursor pagination metadata.
// cursor is the current cursor, nextCursor is the cursor for the next batch.
// If hasMore is true, nextCursor should be the ID of the last item in data.
func CursorPagination[T any](c *gin.Context, items []T, cursor string, nextCursor string, hasMore bool, message string) {
	cp := paginate.NewCursorPagination(cursor, nextCursor, hasMore)
	SuccessWithExtra[[]T, *paginate.CursorPagination](c, items, cp, message)
}
