package app

import (
	"bytes"
	"encoding/json"
	"github.com/strech/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntegrationHandleReversion(t *testing.T) {
	router := Start()
	body := []byte(`{"message":"message to revert - integration"}`)

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
	assert.Equal(t, "message to revert - integration", responseMap["message_to_revert"])
	assert.Equal(t, "integration - revert to message", responseMap["reverted_message"])
}

func TestIntegrationHandleReversion_BadRequestError(t *testing.T) {
	router := Start()
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

func TestIntegrationHandleReversion_InternalServerError(t *testing.T) {
	router := Start()
	body := []byte(`{"message":""}`)

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
