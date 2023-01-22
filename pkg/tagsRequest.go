package pkg

import "sonar/register"

type TagRequest struct {
	Method string
	Path   string
}

func newTagsBuilder() *TagRequest {
	return &TagRequest{}
}

func (r *TagRequest) setMethodType() {
	r.Method = "GET"
}

func (r *TagRequest) setPath() {
	registry := register.GetRegistry()
	r.Path = "/v2/" + registry.Image + "/tags/list"
}

func (r *TagRequest) getRequest() Request {
	return Request{
		Method: r.Method,
		Path:   r.Path,
	}
}
