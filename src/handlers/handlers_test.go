package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProtectedRouteHandler(t *testing.T){
	req, _ := http.NewRequest("GET", "/protected=route", bytes.NewBuffer([]byte(``)))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.ProtectedRoute)

	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("ProtectedRoute handler returned wrong response code: got %d, wanted %d", res.Code, http.StatusOK)
	}
}