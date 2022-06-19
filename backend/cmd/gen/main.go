package main

import (
	"github.com/YReshetko/go-annotation/pkg"

	_ "github.com/YReshetko/go-annotation/annotations/rest"
)

func main() {
	pkg.Process()
}
