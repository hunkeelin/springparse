package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunkeelin/klinutils"
)

type respBody struct {
	Cert         []byte `json:"cert"`
	ChainOfTrust []byte `json:"chainoftrust"`
}

func getcrtv2(g WriteInfo, csrbytes []byte) (*respBody, error) {
	var p respBody
	var dest string
	if g.CA != "" {
		masteraddr, err := klinutils.GetHostnameFromCertv2(g.CA)
		if err != nil {
			return &p, err
		}
		dest = masteraddr
	} else {
		if g.CAName == "" {
			return &p, errors.New("Please specify name of the CA")
		}
		dest = g.CAName
	}
	i := &reqInfo{
		Dest:       dest,
		Dport:      g.CAport,
		Trust:      g.CA,
		TrustBytes: g.CABytes,
		Method:     "POST",
		Headers: map[string]string{
			"content-type": "application/x-www-form-urlencoded",
			"SignCA":       g.SignCA,
		},
		BodyBytes: csrbytes,
		TimeOut:   5000,
	}
	resp, err := sendPayload(i)
	if err != nil {
		return &p, err
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return &p, err
	}
	resp.Body.Close()
	b := body.Bytes()
	if resp.StatusCode != 200 {
		return &p, errors.New(body.String())
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		return &p, err
	}
	return &p, nil
}
func getcrt(g WriteInfo, csrbytes []byte) (*respBody, error) {
	var p respBody
	var dest string
	if g.CA != "" {
		masteraddr, err := klinutils.GetHostnameFromCertv2(g.CA)
		if err != nil {
			return &p, fmt.Errorf("unable to get masteraddr from cert %v", err)
		}
		dest = masteraddr
	} else {
		if g.CAName == "" {
			return &p, errors.New("Please specify name of the CA")
		}
		dest = g.CAName
	}
	i := &reqInfo{
		Dest:       dest,
		Dport:      g.CAport,
		Trust:      g.CA,
		TrustBytes: g.CABytes,
		Method:     "GET",
		Headers: map[string]string{
			"content-type": "application/x-www-form-urlencoded",
			"SignCA":       g.SignCA,
		},
		BodyBytes: csrbytes,
		TimeOut:   5000,
	}
	resp, err := sendPayload(i)
	if err != nil {
		return &p, err
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return &p, fmt.Errorf("unable to readfrom respond body %v", err)
	}
	defer resp.Body.Close()
	b := body.Bytes()
	if resp.StatusCode != 200 {
		return &p, fmt.Errorf("CA server is not giving the cert back")
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		return &p, err
	}
	return &p, nil
}
