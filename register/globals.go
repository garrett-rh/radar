package register

// Global Variable Registry
type register struct {
	TlsNoVerify bool
	Registry    string
	Image       string
	Reference   string
	Digest      string
	UUID        string
}

var Registry *register

func GetRegistry() *register {
	if Registry == nil {
		Registry = &register{}
	}

	return Registry
}
