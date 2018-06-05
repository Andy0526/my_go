package main

import "fmt"

func nonempty(strings []string) []string {
	idx := 0
	for _, s := range strings {
		if s != "" {
			strings[idx] = s
			idx++
		}
	}
	return strings[:idx]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
		}
		out = append(out, s)
	}
	return out
}

func main() {
	//!+main
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
	//!-main
}
