package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate("tree",
		"lishimeng",
		"cmd/tree/main.go", false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
