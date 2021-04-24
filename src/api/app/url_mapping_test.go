package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lucaspichi06/strings-reverter/src/api/controller"
	"github.com/lucaspichi06/strings-reverter/src/api/domain"
	"github.com/strech/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type reversionService struct {
	revert func(request *domain.ReversionRequest) (*domain.ReversionResponse, error)
}

func (r reversionService) Revert(request *domain.ReversionRequest) (*domain.ReversionResponse, error) {
	return r.revert(request)
}

func TestPing(t *testing.T) {
	router := Start()
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fail()
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	jsonResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", string(jsonResp))
}

func TestHandleReversion(t *testing.T) {
	serviceMock := reversionService{
		revert: func(request *domain.ReversionRequest) (*domain.ReversionResponse, error) {
			return &domain.ReversionResponse{
				MessageToRevert: "message to revert",
				RevertedMessage: "revert to message",
			}, nil
		},
	}

	reversionController := controller.NewReversionController(serviceMock)
	router := start(&controllers{
		revert: reversionController,
		status: controller.NewStatusController(),
	})

	body := []byte(`{"message":"message to revert"}`)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/revert_string", bytes.NewBuffer(body))
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	responseMap := make(map[string]interface{})
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "message to revert", responseMap["message_to_revert"])
	assert.Equal(t, "revert to message", responseMap["reverted_message"])
}

func TestHandleReversion_BadRequestError(t *testing.T) {
	serviceMock := reversionService{
		revert: nil,
	}

	reversionController := controller.NewReversionController(serviceMock)
	router := start(&controllers{
		revert: reversionController,
		status: controller.NewStatusController(),
	})

	body := []byte(`{"message":123}`)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/revert_string", bytes.NewBuffer(body))
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	responseMap := make(map[string]interface{})
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, responseMap["message"], "content has invalid format")
}

func TestHandleReversion_InternalServerError(t *testing.T) {
	serviceMock := reversionService{
		revert: func(request *domain.ReversionRequest) (*domain.ReversionResponse, error) {
			return nil, errors.New("test error")
		},
	}

	reversionController := controller.NewReversionController(serviceMock)
	router := start(&controllers{
		revert: reversionController,
		status: controller.NewStatusController(),
	})

	body := []byte(`{"message":"message to revert"}`)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/revert_string", bytes.NewBuffer(body))
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	responseMap := make(map[string]interface{})
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, responseMap["message"], "error while reverting the message")
}
