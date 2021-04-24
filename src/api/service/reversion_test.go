package service

import (
	"github.com/lucaspichi06/strings-reverter/src/api/domain"
	"reflect"
	"testing"
)

func Test_reversion_Revert(t *testing.T) {
	type args struct {
		request *domain.ReversionRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.ReversionResponse
		wantErr bool
	}{
		// Test Uses Cases.
		{"successful reversion", args{request: &domain.ReversionRequest{Message: "test message to revert"}}, &domain.ReversionResponse{
			MessageToRevert: "test message to revert",
			RevertedMessage: "revert to message test",
		}, false},
		{"empty message to revert", args{request: &domain.ReversionRequest{Message: ""}}, nil, true},
		{"invalid characters", args{request: &domain.ReversionRequest{Message: "|@#¢∞¬÷“”≠´‚¡'¿?=)(/&%$·!ªº\"`¨¨ç+æœ€ø®ƒå∫ √å∂å© √µß§"}}, &domain.ReversionResponse{
			MessageToRevert: "|@#¢∞¬÷“”≠´‚¡'¿?=)(/&%$·!ªº\"`¨¨ç+æœ€ø®ƒå∫ √å∂å© √µß§",
			RevertedMessage: "√µß§ √å∂å© |@#¢∞¬÷“”≠´‚¡'¿?=)(/&%$·!ªº\"`¨¨ç+æœ€ø®ƒå∫",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := reversion{}
			got, err := r.Revert(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Revert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Revert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
