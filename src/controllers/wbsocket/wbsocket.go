package wbsocket

import (
	"code.google.com/p/go.net/websocket"
	//zk "fxl.com/utils"
	"log"
)

type Connection *websocket.Conn
type ZkSocket struct {
	cn Connection
}

var zs *ZkSocket = &ZkSocket{}
var c = make(chan bool)

func EchoMessage(ws *websocket.Conn) {
	var content string
	for {
		err := websocket.Message.Receive(ws, &content)
		// If user closes or refreshes the browser, a err will occur
		if err != nil {
			return
		}
		log.Println("here")

	}
	//websocket.Message.Send(ws, zk.GetZooJson("/"))
}
