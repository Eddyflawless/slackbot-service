package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())

}

func TestHttpRequest(t *testing.T) {

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"status\": \"200\"}")
	})

	req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)

	w := httptest.NewRecorder()
	testHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	if resp.StatusCode != 200 {
		t.Fatalf("Expecting status_code of %d but got %d ", 200, resp.StatusCode)
	}

}
