package pkg

type ImageRequest struct {
	Method string
	Path   string
}

func newImageBuilder() *ImageRequest {
	return &ImageRequest{}
}

func (r *ImageRequest) setMethodType() {
	r.Method = "GET"
}

func (r *ImageRequest) setPath() {
	r.Path = "/v2/_catalog"
}

func (r *ImageRequest) getRequest() Request {
	return Request{
		Method: r.Method,
		Path:   r.Path,
	}
}
