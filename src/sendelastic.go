package springparse

import (
	"fmt"
)

func sendElasticSearch(s []byte) error {
	out, err := parseLog(s)
	if err != nil {
		return err
	}
	fmt.Println(s.content, s.id)
	return nil
}
