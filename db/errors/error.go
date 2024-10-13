package errors

import "errors"

var DriverNotSupported = NewFromString("driver not supported")

var _ error = (*DBError)(nil)

type DBError struct {
	Err error
}

func New(err error) error {
	return &DBError{err}
}

func NewFromString(str string) error {
	return New(errors.New(str))
}

func (d *DBError) Error() string {
	return d.Err.Error()
}
