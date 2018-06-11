package main

// +build ignore

import (
	"runtime"
	"fmt"
	"time"
)

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.Compiler)
	time.ANSIC
}
