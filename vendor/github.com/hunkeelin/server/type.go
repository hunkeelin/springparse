package server

import (
	"net/http"
)

// Keycrt the key and cert struct that stores the bytes
type Keycrt struct {
	Cb, Kb []byte
}

// ServerConfig is the configuration for the server
type ServerConfig struct {
	BindAddr     string            // This is the bind address
	BindPort     string            // this is the bind port
	Cert         string            //the location of the .crt for https
	Key          string            // the location of the .key for https
	CertBytes    [][]byte          // the .crt in bytes will take preceding over Cert
	KeyBytes     [][]byte          // the .key in bytes will take preceding over key
	Trust        string            // trust cert location
	TrustBytes   [][]byte          // trust cert in  bytes will take preceding over Trust
	Https        bool              // whether to host in https or not
	Verify       bool              // whether to do http verify or not
	ReadTimeout  int               // read timeout
	WriteTimeout int               // write timeout
	IdleTimeout  int               // idle timeout
	ServeMux     http.Handler      // the http.ServeMux
	Name2cert    map[string]Keycrt // key == hostname, value == cert in bytes
	SNIoverride  bool              // whether to override the sni from name2cert.
}
