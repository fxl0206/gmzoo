package wbsocket

import (
	"code.google.com/p/go.net/websocket"
	zk "zookeeper.fxl.com"
)

func EchoMessage(ws *websocket.Conn) {
	websocket.Message.Send(ws, zk.GetZooJson("/"))
}
