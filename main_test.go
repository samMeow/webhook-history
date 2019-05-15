package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestMainRouteShouldAddQueryStringToRepositroy(t *testing.T) {
	store := []RequestHistory{}
	repo := &StorageImpl{
		Store: store,
	}
	router := NewMainRouter(repo)
	w := performRequest(router, "GET", "/?hello=world")
	assert.Equal(t, http.StatusOK, w.Code)

	s := performRequest(router, "GET", "/history")
	var response []RequestHistory
	err := json.Unmarshal([]byte(s.Body.String()), &response)

	value := response[0].Data
	assert.Nil(t, err)

	assert.Equal(t, "{\"hello\":[\"world\"]}", value)
}
