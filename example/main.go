package main

import (
	"fmt"
	"github.com/zzwx/current"
)

type Type struct {
	current.Path
}

// This example only shows the API usage, however
// actual practical application is in using it in a module containing
// resource files, which the library can access without embedding resources.
func main() {
	r := current.NewPath()
	fmt.Printf("%v\n", r.Path())

	var t Type
	fmt.Printf("%v\n", t.Join("../doc/gobadge.svg"))
}
