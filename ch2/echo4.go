package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "ommit trailing newline") // this creates a pointer
var sep = flag.String("s", " ", "separator")            //this also creates a pointer

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
