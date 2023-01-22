package pkg

// Image request specific version of the Request object
type ImageRequest struct {
	Method string
	Path   string
}

// Returns blank ImageRequest object
func newImageBuilder() *ImageRequest {
	return &ImageRequest{}
}

// Implements only valid HTTP request for this endpoint
func (r *ImageRequest) setMethodType() {
	r.Method = "GET"
}

// Sets path to the image catalog
func (r *ImageRequest) setPath() {
	r.Path = "/v2/_catalog"
}

// Returns ImageRequest object as a generic request
func (r *ImageRequest) getRequest() Request {
	return Request{
		Method: r.Method,
		Path:   r.Path,
	}
}
