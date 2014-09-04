package main

import (
	"conf"
	"gmzoo.com/wbserver"
)

func main() {
	conf.Initail()
	server := wbserver.NewInstance(conf.CurConfFile, conf.CurServerName)
	server.Start()
}
