package springparse

import (
	"testing"
)

func TestValidateOther(t *testing.T) {
	awsElasticSearchURL = "https://foo.bar.com"
	logPrefix = "foo"
	serviceRegex = "foo&bar,candy&bar"
	logDirectory = "/var/log"
	r := New()
	err := r.ValidateOtherEnv()
	if err != nil {
		t.Errorf(err.Error())
	}

}
func TestValidateAws(t *testing.T) {
	r := New()
	err := r.ValidateAwsEnv()
	if err != nil {
		t.Errorf(err.Error())
	}

}
