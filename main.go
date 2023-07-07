package main

import (
	"fmt"

	gomodulesayhello "github.com/dickidarmawansaputra/go-module-sayhello/v2"
)

func main() {
	// fmt.Println(gomodulesayhello.SayHello()) // v1.0.0
	fmt.Println(gomodulesayhello.SayHello("Dicki")) // v2.0.0
}
