package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sonar/register"
	"testing"
)

func TestRootRunner(t *testing.T) {
	expected := "{}"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/v2" && req.URL.Path != "/v2/" {
			t.Errorf("Did not receive request on /v2 or /v2/, recieved on %s", req.URL.Path)
		}
		if req.Method != "GET" {
			t.Errorf("Wrong HTTP Method used. Received %s, want GET", req.Method)
		}

		w.WriteHeader(200)
		_, err := fmt.Fprint(w, expected)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}))
	defer server.Close()

	registry := register.GetRegistry()
	registry.Registry = server.URL

	rootRunner()

}
