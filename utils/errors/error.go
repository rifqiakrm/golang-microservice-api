package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	GetStatus() int
	GetMessage() string
	GetError() string
}

type apiError struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Error string `json:"error,omitempty"`
}

func (e *apiError) GetStatus() int{
	return e.Status
}

func (e *apiError) GetMessage() string{
	return e.Message
}

func (e *apiError) GetError() string{
	return e.Error
}

func NewApiError(statuscode int, message string) ApiError {
	return &apiError{
		Status:  statuscode,
		Message: message,
	}
}

func NewApiErrorFromBytes(body []byte) (ApiError, error) {
	var result apiError

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}

	return &result, nil
}

func NewNotFoundError(message string) ApiError {
	return &apiError{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}