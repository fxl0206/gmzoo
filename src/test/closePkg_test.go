package main

import (
	"log"
	"testing"
)

func Test_Xxx(*testing.T) {
	log.Println("")
	f := pkg()
	f()
}
func main() {
	log.Println("")
	f := pkg()
	f()
}
func pkg() func() {
	type Human struct {
		age  int
		sex  string
		name string
		say  func()
	}
	hum := new(Human)
	i := 1
	hum.age = i
	hum.say = func() {
		//return h.sex

	}
	return func() {
		log.Println(hum)
	}
}
