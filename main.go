package main

import (
	"fmt"
)

func main() {
	var x any = "Era"

	str, ok := x.(string)

	if ok {
		fmt.Println(str, "is a string")
	} else {
		fmt.Println("Not a string")
	}
}
