package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsJSON(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/dog-breeds", nil)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.GetAllDogBreedsJSON)

	handler.ServeHTTP(rr, r)

	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code; got %d, wanted 200", rr.Code)
	}
}
