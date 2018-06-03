package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
