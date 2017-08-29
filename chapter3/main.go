package main

import (
	"fmt"
)

func main() {
	const x string = "Hello, World"
	fmt.Println(x)
}

func f() {
	fmt.Println(x)
}
