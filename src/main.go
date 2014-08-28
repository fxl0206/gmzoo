package main

import (
	"fxl.com/wbserver"
	"os"
)

const (
	defaultConfFile   string = "conf/gmzoo.ini"
	defaultServerName string = "webserver"
)

func main() {
	confPath := defaultConfFile
	serverName := defaultServerName
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
