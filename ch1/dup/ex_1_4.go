// Modify `dup2` to print the names of all files in which each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Need a file name here")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4 %v\n", err)
			}

			input := bufio.NewScanner(f)
			for input.Scan() {
				counts[input.Text()]++
				if names[input.Text()] == nil {
					names[input.Text()] = make(map[string]int)
				} else {
					names[input.Text()][arg]++
				}
			}
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Println("Appeared in: ")
			for name, _ := range names[line] {
				fmt.Printf("%s\n", name)
			}
		}
	}
}
