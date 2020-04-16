package klinpki

type CAConfig struct {
	EmailAddress string
	EcdsaCurve   string
	Certpath     string
	Keypath      string
	MaxDays      float64
	RsaBits      int
	Organization string
}

type SignConfig struct {
	Crtpath  string
	Keypath  string
	CrtBytes []byte
	KeyBytes []byte
	CsrBytes []byte
	Days     float64
	IsCA     bool
}

type CSRConfig struct {
	RsaBits            int
	Country            string
	Province           string
	Locality           string
	EmailAddress       string
	Organization       string
	OrganizationalUnit string
	CommonName         string
	DNSNames           []string
}
