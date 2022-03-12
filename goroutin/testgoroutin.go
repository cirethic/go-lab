package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {

	f("direct")

	go f("goroutine")

	a := func(msg string) int {
		fmt.Println(msg)
		return 1
	}("going")
	fmt.Println(a)

	time.Sleep(time.Second)

	fmt.Println("done")
}
