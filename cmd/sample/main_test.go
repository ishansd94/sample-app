package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ishansd94/sample-app/internal/app/sample"
)

func req(r http.Handler, method, path string, body string) *httptest.ResponseRecorder {

	reader := strings.NewReader(body)

	req, _ := http.NewRequest(method, path, reader)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func TestAPIServer(t *testing.T) {

	apiserver := apiServer()

	// #1
	w := req(apiserver.GetRouter().Handler, "GET", "/version", "")
	if w.Code != http.StatusOK {
		t.Errorf("expected statusOK,  got %v", w.Code)
	}

	// #2
	d, err := json.Marshal(sample.Request{
		Field1: "foo",
		Field2: map[string]string{
			"foo": "bar",
		},
	})
	if err != nil{
		t.Errorf("cannot marshall %v", err)
	}

	w = req(*apiserver.GetRouterHandler(), "POST", "/api/v1/sample", string(d))
	if w.Code != http.StatusCreated {
		t.Errorf("expected statusCreated,  got %v, respones: %v", w.Code, w.Body)
	}

	// #3
	w = req(*apiserver.GetRouterHandler(), "POST", "/api/v1/sample", "")
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected statusBadRequest,  got %v, respones: %v", w.Code, w.Body)
	}

	// #4
	w = req(*apiserver.GetRouterHandler(), "GET", "/api/v1/sample", "")
	if w.Code != http.StatusOK {
		t.Errorf("expected statusOK,  got %v, respones: %v", w.Code, w.Body)
	}
}