package main

import (
	"github.com/hunkeelin/springparse/src"
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
	go springparse.SpringParse()
	err = springparse.Server()
	if err != nil {
		panic(err)
	}
}
