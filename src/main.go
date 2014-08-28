package main

import (
	"fxl.com/wbserver"
	"os"
)

func main() {
	confPath := "./conf/gmzoo.ini"
	serverName := "webserver"
	args := os.Args
	if len(args) > 1 {
		confPath = args[1]
	}
	if len(args) > 2 {
		serverName = args[2]
	}
	server := wbserver.NewInstance(confPath, serverName)
	server.Start()
}
