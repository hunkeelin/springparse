package springparse

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	_, err := ioutil.ReadFile("test.log")
	if err != nil {
		t.Errorf("Unable to read file" + err.Error())
	}
}
func TestSpringparse(t *testing.T) {
	r := New()
	logDirectory = "/tmp"
	serviceRegexList = []string{"banking&service", "foo&bar"}
	r.SpringParse()
}
func TestListDirectory(t *testing.T) {
	logDirectory = "/tmp"
	f, err := listDirectory()
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(f)
}
