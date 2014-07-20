package main

import (
	"fmt"
	"runtime"
)

var ch = make(chan int)

func loop(s string) {
	fmt.Printf("%s started!", s)
	for i := 0; i < 10000; i++ {
		fmt.Printf("%s%d ", s, i)
	}
	ch <- 1
	fmt.Printf("%s finished!", s)
}
func loop2(s string) {
	fmt.Printf("%s started!", s)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s%d ", s, i)
	}
	fmt.Printf("%s finished!", s)

}
func main() {
	runtime.GOMAXPROCS(8)
	go loop("a")
	go loop("c")
	<-ch
	loop2("b")
	<-ch
}
