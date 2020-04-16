package client

import (
	"github.com/hunkeelin/pki"
)

type WriteInfo struct {
	Path      string
	CA        string // CA in file location for trust
	CABytes   []byte // CA in bytes
	CAport    string // CA Ports
	Chain     bool
	CAName    string             // Hostname of the CA. it will default to the hostname from the CA certs if this is left blank
	CSRConfig *klinpki.CSRConfig // import github.com/hunkeelin/pki and read the godocs
	SignCA    string
	Csr       []byte
	Key       []byte
}
