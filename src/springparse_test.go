package springparse

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	_, err := ioutil.ReadFile("test.log")
	if err != nil {
		t.Errorf("Unable to read file" + err.Error())
	}
}
func TestSpringparse(t *testing.T) {
	if os.Getenv("LOG_DIRECTORY") == "" {
		return
	}
	r := NewRunner()
	logDirectory = "/tmp"
	serviceRegexList = []string{"banking&service", "foo&bar"}
	r.SpringParse()
}
func TestListDirectory(t *testing.T) {
	logDirectory = "/tmp"
	r := NewRunner()
	_, err := r.listDirectory()
	if err != nil {
		t.Errorf(err.Error())
	}
}
