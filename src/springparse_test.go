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
	if os.Getenv("LOG_DIRECTORY") == "" || os.Getenv("TAIL_BIN") == "" {
		return
	}
	err := SpringParse()
	if err != nil {
		t.Errorf(err.Error())
	}
}
