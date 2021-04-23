package errors

import (
	"github.com/lucaspichi06/strings-reverter/src/api/domain"
	"net/http"
)

func NewBadRequestAppError(message string) domain.AppError {
	return domain.AppError{message, "bad_request", http.StatusBadRequest, domain.CauseList{}}
}

func NewInternalServerAppError(message string, err error) domain.AppError {
	cause := domain.CauseList{}
	if err != nil {
		cause = append(cause, err.Error())
	}
	return domain.AppError{message, "internal_server_error", http.StatusInternalServerError, cause}
}
