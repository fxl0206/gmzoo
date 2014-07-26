package main

import (
	"fxl.com/wbserver"
)

func main() {
	server := wbserver.WbServer{"8000", "/"}
	server.Start()
}
