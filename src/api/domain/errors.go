package domain

import (
	"fmt"
	"net/http"
)

type CauseList []interface{}

type appError struct {
	ErrorMessage string    `json:"message"`
	ErrorCode    string    `json:"error"`
	ErrorStatus  int       `json:"status"`
	ErrorCause   CauseList `json:"cause"`
}

func (ae appError) Error() string {
	if ae.ErrorCause != nil {
		return fmt.Sprintf("an error of type: %s with value: %s and cause: %v", ae.ErrorCode, ae.ErrorMessage, ae.ErrorCause)
	}

	return fmt.Sprintf("an error of type: %s, with value: %v", ae.ErrorCode, ae.ErrorMessage)
}

func NewBadRequestAppError(message string) appError {
	return appError{message, "bad_request", http.StatusBadRequest, CauseList{}}
}

func NewInternalServerAppError(message string, err error) appError {
	cause := CauseList{}
	if err != nil {
		cause = append(cause, err.Error())
	}
	return appError{message, "internal_server_error", http.StatusInternalServerError, cause}
}
