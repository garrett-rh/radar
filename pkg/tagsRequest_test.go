package pkg

import (
	"reflect"
	"sonar/register"
	"testing"
)

func TestNewTagsBuilder(t *testing.T) {
	got := newTagsBuilder()
	if reflect.TypeOf(&got) == reflect.TypeOf(&TagRequest{}) {
		t.Errorf("Didn't receive type of TagRequest. Recieved: %v", reflect.TypeOf(got))
	}

	if !reflect.DeepEqual(got, &TagRequest{}) {
		t.Errorf("New Tag Request was not empty %v", got)
	}
}

func TestTagSetMethodType(t *testing.T) {
	got := newTagsBuilder()
	got.setMethodType()

	if got.Method != "GET" {
		t.Errorf("Method incorrect, should be GET recieved %s", got.Method)
	}
}

func TestTagSetPath(t *testing.T) {
	registry := register.GetRegistry()
	registry.Image = "blah"
	got := newTagsBuilder()
	got.setPath()

	if got.Path != "/v2/blah/tags/list" {
		t.Errorf("Path incorrect, expected /v2/blah/tags/list, recieved %s", got.Path)
	}
}

func TestTagsGetRequest(t *testing.T) {
	tagsBuilder := newTagsBuilder()
	got := (*tagsBuilder).getRequest()

	if reflect.TypeOf(&got) != reflect.TypeOf(&Request{}) {
		t.Errorf("Wanted type Request, got %v", reflect.TypeOf(got))
	}
}
