package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("Lift off!")
}

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countDown := 10; countDown > 0; countDown-- {
		fmt.Println(countDown)
		<-tick
	}
	launch()
}
