package springparse

import (
	"io/ioutil"
	"testing"
)

func TestPrint(t *testing.T) {
	_, err := ioutil.ReadFile("test.log")
	if err != nil {
		t.Errorf("Unable to read file" + err.Error())
	}
}
func TestSpringparse(t *testing.T) {
	if logDirectory == "" || tailBinary == "" {
		return
	}
	err := springParse()
	if err != nil {
		t.Errorf(err.Error())
	}
}
