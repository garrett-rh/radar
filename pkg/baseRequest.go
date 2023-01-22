package pkg

type BaseRequest struct {
	Method string
	Path   string
}

func newBaseBuilder() *BaseRequest {
	return &BaseRequest{}
}

func (r *BaseRequest) setMethodType() {
	r.Method = "GET"
}

func (r *BaseRequest) setPath() {
	r.Path = "/v2"
}

func (r *BaseRequest) getRequest() Request {
	return Request{
		Method: r.Method,
		Path:   r.Path,
	}
}
