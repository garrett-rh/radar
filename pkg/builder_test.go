package pkg

import (
	"radar/register"
	"testing"
)

func TestGetRequestBuilder(t *testing.T) {
	registry := register.GetRegistry()
	base := GetRequestBuilder("base")
	base.setMethodType()
	base.setPath()

	if base.getRequest().Path != "/v2" && base.getRequest().Path != "/2" {
		t.Errorf("Error creating base request, recieved %v, want %v", base.getRequest(), Request{Method: "GET", Path: "/v2"})
	}

	registry.Image = "blah"
	tags := GetRequestBuilder("tags")
	tags.setMethodType()
	tags.setPath()

	if tags.getRequest().Path != "/v2/blah/tags/list" && tags.getRequest().Method != "GET" {
		t.Errorf("Error creating tags request, recieved %v, want %v", tags.getRequest(), Request{Method: "GET", Path: "/v2/blah/tags/list"})
	}

	image := GetRequestBuilder("image")
	image.setMethodType()
	image.setPath()

	if image.getRequest().Path != "/v2/_catalog" && image.getRequest().Method != "GET" {
		t.Errorf("Error creating iimage request, recieved %v, want %v", image.getRequest(), Request{Method: "GET", Path: "/v2/_catalog"})
	}
}
