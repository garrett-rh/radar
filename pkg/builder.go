package pkg

// Interface that the *Requests will interface with
type RequestBuilder interface {
	setMethodType()
	setPath()
	getRequest() Request
}

// Decides which type of object will be built based on a string passed in by the caller
func GetRequestBuilder(buildType string) RequestBuilder {
	if buildType == "base" {
		return newBaseBuilder()
	}

	if buildType == "tags" {
		return newTagsBuilder()
	}

	if buildType == "image" {
		return newImageBuilder()
	}

	return nil
}
