package pkg

// Director object that implements the RequestBuilder interface
// Assists in creation of Requests
type Director struct {
	builder RequestBuilder
}

// Gives back blank director. Used to wrap the build process
func NewDirector(r RequestBuilder) *Director {
	return &Director{
		builder: r,
	}
}

// Optional call to set the build object to a different object than was passed into the NewDirector call
func (d *Director) SetBuilder(r RequestBuilder) {
	d.builder = r
}

// Builds out request & returns it to the caller
func (d *Director) BuildRequest() Request {

	d.builder.setMethodType()
	d.builder.setPath()

	return d.builder.getRequest()
}
