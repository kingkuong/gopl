package main

import (
	"fmt"
	"os"
)

func main() {
	var s, separator string
	for _, arg := range os.Args[1:] {
		s += separator + arg
		separator = " "
	}

	fmt.Println(s)
}
