package main

import "fmt"
import "flag"
import "strings"

var n = flag.Bool("n", false, "omit trailing newline")
var seq = flag.String("s", " ", "spearator")

func main() {
	flag.Parse()
	fmt.Printf(strings.Join(flag.Args(), *seq))
	if !*n {
		fmt.Println()
	}
}
