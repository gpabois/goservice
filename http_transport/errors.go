package http_transport

import "github.com/gpabois/gostd/serde"

// Convert the error
func HttpError_From(err error) HttpError {
	switch err.(type) {
	case HttpError:
		return err.(HttpError)
	case serde.UnhandledContentType:
		return NewUnsupportedMediaTypeError(err)
	case serde.DeserializeError:
		return NewBadRequestError(err)
	default:
		return NewInternalServerError(err)
	}
}

type HttpError interface {
	Error() string
	Code() int
}

type InternalServerError struct {
	inner error
}

func (err InternalServerError) Error() string {
	return err.inner.Error()
}

func (err InternalServerError) Code() int {
	return 500
}

func NewInternalServerError(err error) HttpError {
	return InternalServerError{inner: err}
}

type NotFoundError struct {
	inner error
}

func (err NotFoundError) Error() string {
	return err.inner.Error()
}

func (err NotFoundError) Code() int {
	return 404
}

func NewNotFoundError(err error) error {
	return NotFoundError{inner: err}
}

type UnsupportedMediaTypeError struct {
	inner error
}

func (err UnsupportedMediaTypeError) Error() string {
	return err.inner.Error()
}

func (err UnsupportedMediaTypeError) Code() int {
	return 415
}

func NewUnsupportedMediaTypeError(err error) HttpError {
	return UnsupportedMediaTypeError{inner: err}
}

type BadRequestError struct {
	inner error
}

func (err BadRequestError) Error() string {
	return err.inner.Error()
}

func (err BadRequestError) Code() int {
	return 400
}

func NewBadRequestError(err error) HttpError {
	return UnsupportedMediaTypeError{inner: err}
}
