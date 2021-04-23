package service

import (
	"errors"
	"fmt"
	"github.com/lucaspichi06/strings-reverter/src/api/domain"
	"strings"
)

type reversion struct{}

func NewReversionService() *reversion {
	return &reversion{}
}

// Revert reverses string using strings.Builder
func (r *reversion) Revert(request *domain.ReversionRequest) (*domain.ReversionResponse, error) {
	if strings.TrimSpace(request.Message) == "" {
		return nil, errors.New("ups... there is nothing to revert")
	}

	var sb strings.Builder
	message := strings.Split(request.Message, " ")
	for i := len(message) - 1; 0 <= i; i-- {
		if _, err := sb.WriteString(fmt.Sprintf("%s ", message[i])); err != nil {
			return nil, err
		}
	}

	return &domain.ReversionResponse{
		MessageToRevert: request.Message,
		RevertedMessage: strings.TrimSpace(sb.String()),
	}, nil
}
