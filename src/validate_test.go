package springparse

import (
	"testing"
)

func TestValidateOther(t *testing.T) {
	awsElasticSearchURL = "https://foo.bar.com"
	logPrefix = "foo"
	serviceRegex = "foo&bar,candy&bar"
	logDirectory = "/var/log"
	err := ValidateOtherEnv()
	if err != nil {
		t.Errorf(err.Error())
	}

}
func TestValidateAws(t *testing.T) {
	err := ValidateAwsEnv()
	if err != nil {
		t.Errorf(err.Error())
	}

}
