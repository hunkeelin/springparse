package klinutils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/theckman/go-flock"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

func Dowork() {
	fmt.Println("start")
	time.Sleep(9 * time.Second)
	fmt.Println("end")
}
func Is_mac(s string) bool {
	match, _ := regexp.MatchString("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$", s)
	return match
}

func ThreadRunner(n int) chan func() {
	fc := make(chan func())
	for i := 0; i < n; i++ {
		go func() {
			for {
				f := <-fc
				f()
			}
		}()
	}
	return fc
}

// Example
/*
func contest() {
    f := runner(5)
    var wg sync.WaitGroup
    for i := 0; i < 15; i++ {
        wg.Add(1)
        j := i
        f <- func() {
            Dowork()
            wg.Done()
        }
    }
    wg.Wait()
}
*/
func Is_ipv4(host string) bool {
	parts := strings.Split(host, ".")
	if len(parts) < 4 {
		return false
	}
	for _, x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func verifySignature(secret []byte, signature string, body []byte) bool {
	const signaturePrefix = "sha1="
	const signatureLength = 45 // len(SignaturePrefix) + len(hex(sha1))
	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}
	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))
	return hmac.Equal(signBody(secret, body), actual)
}

func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}

	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return hostname
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip, err := ipv4.MarshalText()
			if err != nil {
				return hostname
			}
			hosts, err := net.LookupAddr(string(ip))
			if err != nil || len(hosts) == 0 {
				return hostname
			}
			fqdn := hosts[0]
			return strings.TrimSuffix(fqdn, ".")
		}
	}
	return hostname
}
func Joblist(path string) map[string]string {
	m := make(map[string]string)
	err := os.Chdir(path)
	if err != nil {
		log.Fatal("no such file or directory ", path)
	}
	ls, _ := filepath.Glob("*")
	for _, provider := range ls { //github.com and other provider
		os.Chdir(path + provider)
		orgs, _ := filepath.Glob("*")
		for _, org := range orgs { //github.com/orgs
			os.Chdir(path + provider + "/" + org)
			jobs, _ := filepath.Glob("*")
			for _, job := range jobs { // going through each jobs
				confdir, _ := os.Getwd()
				confdir = confdir + "/" + job + "/" + "config.xml"
				url := "https://" + provider + "/" + org + "/" + job
				m[strings.ToLower(url)] = confdir

			}
		}
	}
	return m
}

func GetHostnameFromCertv2(path string) (string, error) {
	var s string
	var err error
	e, err := ioutil.ReadFile(path)
	if err != nil {
		return s, err
	}
	block, _ := pem.Decode(e)
	if block == nil {
		return s, errors.New("Unable to decode the cert file")
	}
	leaf, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return s, err
	}
	return leaf.DNSNames[0], nil
}

func Matchstring(s, regex string) bool {
	match, err := regexp.MatchString(regex, s)
	if err != nil {
		log.Fatal("regex matching problem")
	}
	return match
}
func Waitforqueue(dir string) *flock.Flock {
	fileLock := flock.NewFlock(dir)

	locked, err := fileLock.TryLock()

	if err != nil {
		log.Fatal("unable to lock the file at ", dir)
	}

	if locked {
		return fileLock
	} else {
		fmt.Println("wait 10 seconds there's another process running")
		time.Sleep(10 * time.Second)
		return Waitforqueue(dir)
	}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Exist(p string) bool {
	if _, err := os.Stat(p); err == nil {
		return true
	}
	return false
}
func Runshell(cmd string, args []string, uid, gid uint32) error {
	//err := exec.Command(cmd, args...).Run()
	acmd := exec.Command(cmd, args...)
	acmd.SysProcAttr = &syscall.SysProcAttr{}
	acmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}
	out, err := acmd.CombinedOutput()
	fmt.Printf("%s\n", out)
	return err
}
func Runshellv2(cmd string, args []string) error {
	err := exec.Command(cmd, args...).Run()
	return err
}

func Removestring(s []string, pattern string) []string {
	var toreturn []string
	for _, raw_element := range s {
		element := strings.Replace(raw_element, " ", "", -1)
		if strings.HasPrefix(element, pattern) {
			ele := strings.Replace(element, pattern, "", -1)
			if strings.HasPrefix(ele, "HEAD") || strings.HasPrefix(ele, "master") {
				continue
			}
			toreturn = append(toreturn, ele)
		}
	}
	return toreturn
}

func Outshell(cmd string, args []string) (string, error) {
	output, err := exec.Command(cmd, args...).Output()
	return string(output), err
}
func Cleandir(s, env []string, workers int) {
	sema := make(chan struct{}, workers)
	wg := sync.WaitGroup{}
	for _, element := range s {
		if !StringInSlice(element, env) {
			sema <- struct{}{}
			wg.Add(1)
			go func(element string) {
				os.RemoveAll(element)
				<-sema
				wg.Done()
			}(element)
		}
	}
	wg.Wait()
}

func Createdir(s, env []string, workers int) {
	sema := make(chan struct{}, workers)
	wg := sync.WaitGroup{}
	for _, element := range env {
		if StringInSlice(element, s) == false {
			sema <- struct{}{}
			wg.Add(1)
			go func(element string) {
				os.MkdirAll(element+"/"+"modules", 0755)
				<-sema
				wg.Done()
			}(element)
		}
	}
	wg.Wait()
}
func Checkstring(s, pattern string) {
	if strings.HasPrefix(s, "mod") {
		if string(s[len(s)-1]) != "," {
			log.Fatal("error missing comma on line: ", s)
		}
	}

	checknum := 0

	for _, r := range s {
		c := string(r)
		if c == "'" {
			checknum = checknum + 1
		}
	}

	if checknum != 2 {
		log.Fatal("error missing single quotes on line: ", s)
	}
}
func Checkerr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Trim(x string) string {
	pattern := "'"
	Checkstring(x, pattern)
	bra := strings.Index(x, pattern)
	if bra < 0 {
		return ""
	}
	rx := x[bra+1:]
	ket := strings.Index(rx, pattern)
	return rx[:ket]
}
func signBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}
func VerifySignature(secret []byte, signature string, body []byte) bool {

	const signaturePrefix = "sha1="
	const signatureLength = 45 // len(SignaturePrefix) + len(hex(sha1))

	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))

	return hmac.Equal(signBody(secret, body), actual)
}

func Isvalidmethod(r *http.Request) bool {
	methodlist := []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTOINS", "TRACE"}
	return StringInSlice(r.Method, methodlist)
}
