package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	HelloHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if string(data) != "Hello MIFI!" {
		t.Errorf("wrong response")
	}
}
