package errors

import "net/http"

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Unauthorized() error {
	return &HttpError{Code: http.StatusUnauthorized, Message: "unauthorized"}
}

func Forbidden() error {
	return &HttpError{Code: http.StatusForbidden, Message: "forbidden"}
}

func BadRequest() error {
	return &HttpError{Code: http.StatusBadRequest, Message: "bad request"}
}

func BadGateway() error {
	return &HttpError{Code: http.StatusBadGateway, Message: "bad gateway"}
}

func (h *HttpError) Error() string {
	return h.Message
}
