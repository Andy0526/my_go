package main

import (
	"os"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	for i := 1; i <= len(os.Args); i++ {
		fmt.Printf(" %s\n", comma(os.Args[i]))
	}
}
