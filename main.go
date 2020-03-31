package main

import (
	"github.com/hunkeelin/springparse/src"
	"time"
)

func main() {
	err := springparse.ValidateAwsEnv()
	if err != nil {
		panic(err)
	}
	err = springparse.ValidateOtherEnv()
	if err != nil {
		panic(err)
	}
	springparse.ShowConfiguration()
	r := springparse.New()
	go func() {
		for {
			r.SpringParse()
			time.Sleep(3 * time.Second)
		}
	}()
	err = springparse.Server()
	if err != nil {
		panic(err)
	}
}
