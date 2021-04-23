package domain

import "fmt"

type CauseList []interface{}

type AppError struct {
	ErrorMessage string    `json:"message"`
	ErrorCode    string    `json:"error"`
	ErrorStatus  int       `json:"status"`
	ErrorCause   CauseList `json:"cause"`
}

func (ae AppError) Error() string {
	if ae.ErrorCause != nil {
		return fmt.Sprintf("an error of type: %s with value: %s and cause: %v", ae.ErrorCode, ae.ErrorMessage, ae.ErrorCause)
	}

	return fmt.Sprintf("an error of type: %s, with value: %v", ae.ErrorCode, ae.ErrorMessage)
}