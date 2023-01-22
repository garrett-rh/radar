package pkg

import (
	"reflect"
	"testing"
)

func TestNewImageBuilder(t *testing.T) {
	got := newImageBuilder()
	if reflect.TypeOf(&got) == reflect.TypeOf(&ImageRequest{}) {
		t.Errorf("Didn't receive type of ImageRequest. Recieved: %v", reflect.TypeOf(got))
	}

	if !reflect.DeepEqual(got, &ImageRequest{}) {
		t.Errorf("New Image Request was not empty %v", got)
	}
}

func TestImageSetMethodType(t *testing.T) {
	got := newImageBuilder()
	got.setMethodType()

	if got.Method != "GET" {
		t.Errorf("Method incorrect, should be GET recieved %s", got.Method)
	}
}

func TestImageSetPath(t *testing.T) {
	got := newImageBuilder()
	got.setPath()

	if got.Path != "/v2/_catalog" {
		t.Errorf("Path incorrect, expected /v2, recieved %s", got.Path)
	}
}

func TestImageGetRequest(t *testing.T) {
	imageBuilder := newImageBuilder()
	got := (*imageBuilder).getRequest()

	if reflect.TypeOf(&got) != reflect.TypeOf(&Request{}) {
		t.Errorf("Wanted type Request, got %v", reflect.TypeOf(got))
	}
}
