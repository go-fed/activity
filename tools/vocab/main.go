package main

import (
	"github.com/go-fed/activity/tools/defs"
	"github.com/go-fed/activity/tools/vocab/gen"
	"io/ioutil"
)

func main() {
	allTypes := append(defs.AllCoreTypes, defs.AllExtendedTypes...)
	b, err := gen.GenerateImplementations(allTypes, defs.AllPropertyTypes, defs.AllValueTypes)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("vocab.go", b, 0666)
	if err != nil {
		panic(err)
	}
}
