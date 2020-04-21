package main

import (
	"github.com/hunkeelin/springparse"
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
	//initialize the elasticsearch client
	r.InitElastic()
	r.SpringParse()
	err = r.Server()
	if err != nil {
		panic(err)
	}
}
