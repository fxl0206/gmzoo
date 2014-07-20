package main

import (
	"fmt"
	"runtime"
)

var ii = 0

func main() {
	runtime.GOMAXPROCS(8)
	ch := make(chan int)
	task("A", ch)
	task("B", ch)
	fmt.Printf("begin\n")
	<-ch
}
func task(name string, ch chan int) {
	go func() {
		for ii < 51 {
			fmt.Printf("%s %d\n", name, ii)
			ii++
			runtime.Gosched()
			if ii == 50 {
				ch <- 1
				break
			}
		}
	}()
}
