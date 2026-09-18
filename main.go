package main

import (
	"fmt"
)

func main() {
	name := "e"
	if n := len(name); n == 0 {
		fmt.Println("No name provided")
	} else {
		fmt.Println("Hello", name)
	}
}
