// +build http

package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MyTestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadGateway)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"active" : true}`)
}

func TestMyHttp(t *testing.T) {
	req, err := http.NewRequest("GET", "/handlerTest", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		w.Header().Set("Content-Type", "application/json")

		io.WriteString(w, `{"active" : true}`)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("handler response statuscode is not Ok,status :", status)
	}

	expected := `{"active" : true}`
	if rr.Body.String() != expected {
		t.Error("handler response body is not ", expected, ",body :", rr.Body.String())
	}

}
