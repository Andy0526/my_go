package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for idx, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(idx, arg)
	}
	fmt.Println(s)
}
