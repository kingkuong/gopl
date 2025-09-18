package main

import (
	"fmt"
	"gopl/ch2/popcount"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	t, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		fmt.Println("Errored convering")
	}

	fmt.Println("Counting....")
	fmt.Printf("%s pop count is %d\n", arg, popcount.PopCount(t))
}
