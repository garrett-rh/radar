package register

// Global Variable Registry
type register struct {
	// Set via the -k or --insecure flag to disable cert validation checks
	TlsNoVerify bool
	// URI of the registry
	Registry string
	// Image name passed in via radar image $IMAGE_NAME command
	Image string
	// Not yet implemented
	Reference string
	// Not yet implemented
	Digest string
	// Not yet implemented
	UUID string
}

var Registry *register

// Returns global registry
func GetRegistry() *register {
	if Registry == nil {
		Registry = &register{}
	}

	return Registry
}
