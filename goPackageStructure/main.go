package main

import (
	"fmt"

	hello "example.com/hello"
)

func main() {
	helloValue := hello.Hello()
	fmt.Println(helloValue)
}
