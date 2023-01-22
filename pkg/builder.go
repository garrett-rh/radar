package pkg

type RequestBuilder interface {
	setMethodType()
	setPath()
	getRequest() Request
}

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
