package main

import (
	"github.com/widuu"
	"log"
	"testing"
)

func Test(t *testing.T) {
	conf := goini.SetConfig("./conf/gmzoo.ini")
	wbPort := conf.GetValue("webserver", "port")
	log.Println(wbPort) //root
	data := conf.ReadList()
	log.Println(data)
}
