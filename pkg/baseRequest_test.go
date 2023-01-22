package pkg

import (
	"reflect"
	"testing"
)

func TestNewBaseBuilder(t *testing.T) {
	got := newBaseBuilder()
	if reflect.TypeOf(&got) == reflect.TypeOf(&BaseRequest{}) {
		t.Errorf("Didn't receive type of BaseRequest. Recieved: %v", reflect.TypeOf(got))
	}

	if !reflect.DeepEqual(got, &BaseRequest{}) {
		t.Errorf("New Base Request was not empty %v", got)
	}
}

func TestBaseSetMethodType(t *testing.T) {
	got := newBaseBuilder()
	got.setMethodType()

	if got.Method != "GET" {
		t.Errorf("Method incorrect, should be GET recieved %s", got.Method)
	}
}

func TestBaseSetPath(t *testing.T) {
	got := newBaseBuilder()
	got.setPath()

	if got.Path != "/v2" {
		t.Errorf("Path incorrect, expected /v2, recieved %s", got.Path)
	}
}

func TestBaseGetRequest(t *testing.T) {
	baseBuilder := newBaseBuilder()
	got := (*baseBuilder).getRequest()

	if reflect.TypeOf(&got) != reflect.TypeOf(&Request{}) {
		t.Errorf("Wanted type Request, got %v", reflect.TypeOf(got))
	}
}
