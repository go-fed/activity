package main

import (
	"github.com/go-fed/activity/tools/defs"
	"github.com/go-fed/activity/tools/vocab/gen"
	"io/ioutil"
)

func main() {
	allTypes := append(defs.AllCoreTypes, defs.AllExtendedTypes...)
	files, err := gen.GenerateImplementations(allTypes, defs.AllPropertyTypes, defs.AllValueTypes)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		err = ioutil.WriteFile(f.Name, f.Content, 0666)
		if err != nil {
			panic(err)
		}
	}
}
