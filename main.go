package main

import (
	"github.com/hunkeelin/springparse/src"
	"time"
)

func main() {
	r := springparse.NewRunner()
	err := r.ValidateAwsEnv()
	if err != nil {
		panic(err)
	}
	err = r.ValidateOtherEnv()
	if err != nil {
		panic(err)
	}
	r.ShowConfiguration()
	go func() {
		for {
			r.SpringParse()
			time.Sleep(3 * time.Second)
		}
	}()
	err = r.Server()
	if err != nil {
		panic(err)
	}
}
