package client

import (
	"bytes"
	"encoding/pem"
	"errors"
	"github.com/hunkeelin/pki"
)

func Getkeycrtbyte(w WriteInfo) (crtpem, keypem []byte, err error) {
	var crt [][]byte
	var bcrt, bkey bytes.Buffer
	csr, key := klinpki.GenCSRv2(w.CSRConfig)
	if w.CA == "" && len(w.CABytes) == 0 {
		return bcrt.Bytes(), bkey.Bytes(), errors.New("Please specify CA in bytes or give the location of the CA")
	}
	f, err := getcrt(w, csr.Bytes)
	if err != nil {
		return bcrt.Bytes(), bkey.Bytes(), err
	}
	crt = append(crt, f.Cert)
	if w.Chain {
		p, rest := pem.Decode(f.ChainOfTrust)
		crt = append(crt, p.Bytes)
		for len(rest) != 0 {
			p, rest = pem.Decode(rest)
			crt = append(crt, p.Bytes)
		}
	}
	for _, c := range crt {
		pem.Encode(&bcrt, &pem.Block{Type: "CERTIFICATE", Bytes: c})
	}
	pem.Encode(&bkey, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: key.Bytes})
	return bcrt.Bytes(), bkey.Bytes(), nil
}
