package pkg

type Director struct {
	builder RequestBuilder
}

func NewDirector(r RequestBuilder) *Director {
	return &Director{
		builder: r,
	}
}

func (d *Director) SetBuilder(r RequestBuilder) {
	d.builder = r
}

func (d *Director) BuildRequest() Request {

	d.builder.setMethodType()
	d.builder.setPath()

	return d.builder.getRequest()
}
