package main

import (
	"github.com/go-fed/activity/tools/defs"
	"github.com/go-fed/activity/tools/streams/gen"
	"io/ioutil"
)

func main() {
	allTypes := append(defs.AllCoreTypes, defs.AllExtendedTypes...)
	b, err := gen.GenerateConvenienceTypes(allTypes)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("activitystreams.go", b, 0666)
	if err != nil {
		panic(err)
	}
}
