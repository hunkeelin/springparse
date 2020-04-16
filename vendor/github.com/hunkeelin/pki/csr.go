package klinpki

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
)

func GenCSRv2(c *CSRConfig) (*pem.Block, *pem.Block) {
	var hname string
	if c.CommonName == "" {
		hname = Hostname()
	} else {
		hname = c.CommonName
	}

	priv, err := rsa.GenerateKey(rand.Reader, c.RsaBits)
	if err != nil {
		panic(err)
	}
	subj := pkix.Name{
		CommonName:         hname,
		Country:            []string{c.Country},
		Province:           []string{c.Province},
		Locality:           []string{c.Locality},
		OrganizationalUnit: []string{c.OrganizationalUnit},
		Organization:       []string{c.Organization},
	}
	template := x509.CertificateRequest{
		Subject:            subj,
		SignatureAlgorithm: x509.SHA256WithRSA,
		DNSNames:           c.DNSNames,
	}
	template.DNSNames = append(template.DNSNames, hname)
	template.EmailAddresses = append(template.EmailAddresses, c.EmailAddress)
	csrBytes, _ := x509.CreateCertificateRequest(rand.Reader, &template, priv)
	return &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes}, pemBlockForKey(priv)
}
