package server

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPrint(t *testing.T) {
	_, err := ioutil.ReadFile("test.log")
	if err != nil {
		t.Errorf("Unable to read file" + err.Error())
	}
}
