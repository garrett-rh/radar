package pkg

import "radar/register"

type TagRequest struct {
	Method string
	Path   string
}

// Returns an empty TagRequest object
func newTagsBuilder() *TagRequest {
	return &TagRequest{}
}

// Implements only valid HTTP Request for this endpoint
func (r *TagRequest) setMethodType() {
	r.Method = "GET"
}

// Grabs image from the global register & creates the path for listing tags
func (r *TagRequest) setPath() {
	registry := register.GetRegistry()
	r.Path = "/v2/" + registry.Image + "/tags/list"
}

// Returns TagRequest as generic request object
func (r *TagRequest) getRequest() Request {
	return Request{
		Method: r.Method,
		Path:   r.Path,
	}
}
