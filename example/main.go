package main

import (
	"fmt"
	"github.com/zzwx/current"
)

func main() {
	r := current.NewPath()
	s := r.Path()
	fmt.Printf("Result main(): %v\n", s)
}
