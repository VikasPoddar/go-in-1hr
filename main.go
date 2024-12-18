package main

import (
	"fmt"
	"github.com/VikasPoddar/go-in-1hr/os_package"
)

func main() {
	// main package act entry-point to execution
	// every demo has its own package
	// for example
	// 1. testing_demo package: This package demonstrate how simple testing works in go
	os_package.IoDemo()
	os_package.OsDemo()
}

//following function are used for testing github-action's CI workflow

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s\n", name)
}
