package springparse

import (
	"testing"
)

func TestShouldwatch(t *testing.T) {
	serviceRegexList = []string{"banking&service", "foo&bar"}
	r := New()
	s := r.shouldWatch(shouldWatchInput{
		logFile: "/var/log/banking_api_a84dljf_asdl_service.log",
	})
	if !s.watch {
		t.Errorf("Did not match")
	}
}
