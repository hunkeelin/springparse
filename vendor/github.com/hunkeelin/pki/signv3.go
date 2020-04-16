package klinpki

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

func SignCSRv3(s *SignConfig) ([]byte, error) {
	// load CA key pair
	//      public key
	var clientCRTRaw []byte
	if s.CrtBytes == nil || s.KeyBytes == nil {
		return clientCRTRaw, fmt.Errorf("key/certs is nil")
	}
	pemBlock, _ := pem.Decode(s.CrtBytes)
	if pemBlock == nil {
		panic("pem.Decode failed")
	}
	caCRT, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return clientCRTRaw, err
	}

	//      private key
	pemBlock, _ = pem.Decode(s.KeyBytes)
	if pemBlock == nil {
		panic("pem.Decode failed")
	}
	der := pemBlock.Bytes
	//	der, err := x509.DecryptPEMBlock(pemBlock, []byte("shit"))
	//	if err != nil {
	//		return clientCRTRaw,err
	//	}
	caPrivateKey, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		return clientCRTRaw, err
	}

	// load client certificate request
	if err != nil {
		return clientCRTRaw, err
	}
	clientCSR, err := x509.ParseCertificateRequest(s.CsrBytes)
	if err != nil {
		return clientCRTRaw, err
	}
	if err = clientCSR.CheckSignature(); err != nil {
		return clientCRTRaw, err
	}

	// create client certificate template
	clientCRTTemplate := x509.Certificate{
		Signature:          clientCSR.Signature,
		SignatureAlgorithm: clientCSR.SignatureAlgorithm,

		PublicKeyAlgorithm: clientCSR.PublicKeyAlgorithm,
		PublicKey:          clientCSR.PublicKey,

		SerialNumber:          big.NewInt(2),
		Issuer:                caCRT.Subject,
		Subject:               clientCSR.Subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Duration(s.Days*24) * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		DNSNames:              clientCSR.DNSNames,
		EmailAddresses:        clientCSR.EmailAddresses,
		BasicConstraintsValid: s.IsCA,
		IsCA:                  s.IsCA,
	}
	if s.IsCA {
		clientCRTTemplate.BasicConstraintsValid = s.IsCA
		clientCRTTemplate.IsCA = s.IsCA
		clientCRTTemplate.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}
		clientCRTTemplate.KeyUsage |= x509.KeyUsageCertSign
	}

	// create client certificate from template and CA public key
	clientCRTRaw, err = x509.CreateCertificate(rand.Reader, &clientCRTTemplate, caCRT, clientCSR.PublicKey, caPrivateKey)
	if err != nil {
		return clientCRTRaw, err
	}

	// save the certificate
	return clientCRTRaw, nil
}
