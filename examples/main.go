package main

import (
	"github.com/hunkeelin/springparse"
	"time"
)

func main() {
	r := springparse.New()
	err := r.ValidateAwsEnv()
	if err != nil {
		panic(err)
	}
	err = r.ValidateOtherEnv()
	if err != nil {
		panic(err)
	}
	r.ShowConfiguration()
	// host the metric server
	go r.Server()
	for {
		r.SpringParse()
		time.Sleep(3 * time.Second)
	}
}
