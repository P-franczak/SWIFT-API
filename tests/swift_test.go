package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/P-franczak/swift-api/handlers"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateSwiftCode(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/v1/swift-codes", handlers.CreateSwiftCode).Methods("POST")

	payload := `{
		"swiftCode": "TESTPLXXX",
		"bankName": "Test Bank",
		"address": "Test Street 123",
		"countryISO2": "PL",
		"countryName": "POLAND",
		"isHeadquarter": true
	}`

	req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	var result map[string]string
	json.Unmarshal(response.Body.Bytes(), &result)
	assert.Equal(t, "SWIFT code added successfully", result["message"])
}

func TestGetNonExistingSwiftCode(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/v1/swift-codes/{swiftCode}", handlers.GetSwiftCode).Methods("GET")

	req, _ := http.NewRequest("GET", "/v1/swift-codes/INVALID", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusNotFound, response.Code)
}
