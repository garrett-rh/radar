package register

type endpoints struct {
	Base           string
	Catalog        string
	Tags           string
	Manifest       string
	Blob           string
	InitBlobUpload string
	BlobUpload     string
}

var Endpoint *endpoints

func GetEndpoints() *endpoints {
	if Endpoint == nil {
		Endpoint = &endpoints{
			Base:           "/v2",
			Catalog:        "/v2/_catalog",
			Tags:           "/v2/{{.Image}}/tags/list",
			Manifest:       "/v2/{{.Image}}/manifests/{{.Reference}}",
			Blob:           "/v2/{{.Image}}/blobs/{{.Digest}}",
			InitBlobUpload: "/v2/{{.Image}}/blobs/uploads",
			BlobUpload:     "/v2/{{.Image}}/blobs/{{.UUID}}",
		}
	}
	return Endpoint
}
