package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"radar/register"
	"testing"
)

func TestImageRunner(t *testing.T) {
	expected := `{"repositories":[]}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/v2/_catalog" && req.URL.Path != "/v2/_catalog/" {
			t.Errorf("Did not receive request on /v2/_catalog or /v2/_catalog/, recieved on %s", req.URL.Path)
		}
		if req.Method != "GET" {
			t.Errorf("Wrong HTTP Method used. Received %s, want GET", req.Method)
		}
		_, err := fmt.Fprint(w, expected)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}))
	defer server.Close()

	registry := register.GetRegistry()
	registry.Registry = server.URL

	imageRunner([]string{})

}

func TestTagRunner(t *testing.T) {
	expected := `{"name":"blah","tags":["1234"]}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/v2/blah/tags/list" && req.URL.Path != "/v2/blah/tags/list/" {
			t.Errorf("Did not receive request on /v2/_catalog or /v2/_catalog/, recieved on %s", req.URL.Path)
		}
		if req.Method != "GET" {
			t.Errorf("Wrong HTTP Method used. Received %s, want GET", req.Method)
		}

		_, err := fmt.Fprint(w, expected)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}))
	defer server.Close()

	registry := register.GetRegistry()
	registry.Registry = server.URL

	imageRunner([]string{"blah"})

}
