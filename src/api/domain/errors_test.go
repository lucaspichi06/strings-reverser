package domain

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestNewBadRequestAppError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want appError
	}{
		// Test Uses Cases.
		{"bad request error", args{message: "content has an invalid format"}, appError{
			ErrorMessage: "content has an invalid format",
			ErrorCode:    "bad_request",
			ErrorStatus:  http.StatusBadRequest,
			ErrorCause:   CauseList{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBadRequestAppError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBadRequestAppError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInternalServerAppError(t *testing.T) {
	type args struct {
		message string
		err     error
	}
	tests := []struct {
		name string
		args args
		want appError
	}{
		// Test Uses Cases.
		{"internal server error without cause", args{message: "internal server error"}, appError{
			ErrorMessage: "internal server error",
			ErrorCode:    "internal_server_error",
			ErrorStatus:  http.StatusInternalServerError,
			ErrorCause:   CauseList{},
		}},
		{"internal server error with cause", args{message: "internal server error", err: errors.New("test error")}, appError{
			ErrorMessage: "internal server error",
			ErrorCode:    "internal_server_error",
			ErrorStatus:  http.StatusInternalServerError,
			ErrorCause: CauseList{
				"test error",
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInternalServerAppError(tt.args.message, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInternalServerAppError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appError_Error(t *testing.T) {
	type fields struct {
		ErrorMessage string
		ErrorCode    string
		ErrorStatus  int
		ErrorCause   CauseList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// Test Uses Cases.
		{"error with cause", fields{
			ErrorMessage: "internal server error",
			ErrorCode:    "internal_server_error",
			ErrorStatus:  http.StatusInternalServerError,
			ErrorCause: CauseList{
				"test error",
			},
		}, "an error of type: internal_server_error with value: internal server error and cause: [test error]"},
		{"error without cause", fields{
			ErrorMessage: "internal server error",
			ErrorCode:    "internal_server_error",
			ErrorStatus:  http.StatusInternalServerError,
		}, "an error of type: internal_server_error, with value: internal server error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ae := appError{
				ErrorMessage: tt.fields.ErrorMessage,
				ErrorCode:    tt.fields.ErrorCode,
				ErrorStatus:  tt.fields.ErrorStatus,
				ErrorCause:   tt.fields.ErrorCause,
			}
			if got := ae.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}