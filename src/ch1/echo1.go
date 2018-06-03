package main

import (
	"os"
	"fmt"
)

func main() {
	var s, seq string
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
		s += seq + os.Args[i]
		seq = " "
	}
	fmt.Println(s)
}
