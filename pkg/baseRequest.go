package pkg

type BaseRequest struct {
	Method string
	Path   string
}

// Returns a new BaseRequest object
func newBaseBuilder() *BaseRequest {
	return &BaseRequest{}
}

// Sets the BaseRequest to "GET". Only valid HTTP Request for this endpoint
func (r *BaseRequest) setMethodType() {
	r.Method = "GET"
}

// Sets Path to base of registry
func (r *BaseRequest) setPath() {
	r.Path = "/v2"
}

// Translates the base request into a generic Request object
func (r *BaseRequest) getRequest() Request {
	return Request{
		Method: r.Method,
		Path:   r.Path,
	}
}
