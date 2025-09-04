package main

import (
	"fmt"
	"os"
)

func main() {
	var s, seperator string
	for i := 0; i < len(os.Args); i++ {
		s += seperator + os.Args[i]
		seperator = " "
	}

	fmt.Println(s)
}
