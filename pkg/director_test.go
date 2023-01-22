package pkg

import (
	"reflect"
	"testing"
)

func TestNewDirector(t *testing.T) {
	request := GetRequestBuilder("base")
	director := NewDirector(request)

	if reflect.TypeOf(*director) != reflect.TypeOf(Director{builder: request}) {
		t.Errorf("Expected type of Director, got %v", reflect.TypeOf(*director))
	}
}

func TestBuildRequest(t *testing.T) {
	requestBuilder := GetRequestBuilder("base")
	director := NewDirector(requestBuilder)
	request := director.BuildRequest()

	if request.Method != "GET" && request.Path != "/v2/_catalog" {
		t.Errorf("Director did not set new build request. Wanted {GET /v2/_catalog}, recieved %v", request)
	}
}

func TestSetBuilder(t *testing.T) {
	// Tests if the request being built is switched even though initially init
	// with a different request type
	baseRequest := GetRequestBuilder("base")
	imageRequest := GetRequestBuilder("image")

	director := NewDirector(baseRequest)
	director.SetBuilder(imageRequest)

	latestRequest := director.BuildRequest()

	if latestRequest.Method != "GET" && latestRequest.Path != "/v2/_catalog" {
		t.Errorf("Director did not set new build request. Wanted %v, recieved %v", baseRequest, latestRequest)
	}
}
