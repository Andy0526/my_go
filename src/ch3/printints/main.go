package main

import (
	"bytes"
	"fmt"
)

func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for idx, v := range values {
		if idx > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intToString([]int{1, 2, 4}))
}
