package springparse

import (
	"testing"
)

func TestGetcontainerinfo(t *testing.T) {
	f := "/var/log/containers/foo-649df48fdb-7bd9l_api_service-e93ad848d21a967c6951f0bc0ead003ed2b7fcdef9c43c30711d6cb90bb45d2e.log"
	result, err := getContainerInfo(f)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result.podName != "foo-649df48fdb-7bd9l" {
		t.Errorf("podname is not correct")
	}
	if result.nameSpace != "api" {
		t.Errorf("api is not correct")
	}
	if result.containerName != "service" {
		t.Errorf("containerName is not correct")
	}
	if result.dockerId != "e93ad848d21a967c6951f0bc0ead003ed2b7fcdef9c43c30711d6cb90bb45d2e" {
		t.Errorf("dockerId is not correct")
	}
}
