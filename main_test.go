package main

import (
	"net/http/httptest"
    "testing"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
)



func TestCreateEndpoint(t *testing.T) {
	var jsonData = []byte(`{
	"User_id": "1245"
	"Email": "jayavignesh@gmail.com",
	"Password": "password"
	}`)
    request, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonData))
    response := httptest.NewRecorder()
	createUser(response,request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestCreateEndpoint(t *testing.T) {
    request, _ := http.NewRequest("POST", "/api/users/123")
    response := httptest.NewRecorder()
	getUser(response,request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
