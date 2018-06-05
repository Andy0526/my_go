package main

import "fmt"

func appendslice(x []int, y ...int) []int {
	var z []int
	xlen := len(x)
	zlen := +len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap <= 2*xlen {
			zcap = 2 * xlen
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[xlen:], y)
	return z
}

func appendInt(x []int, y int) []int {
	var z []int
	xlen := len(x)
	zlen := xlen + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap <= 2*xlen {
			zcap = 2 * xlen
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[xlen] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
