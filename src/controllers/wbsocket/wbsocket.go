package wbsocket

import (
	"code.google.com/p/go.net/websocket"
	zk "fxl.com/utils"
)

func EchoMessage(ws *websocket.Conn) {
	websocket.Message.Send(ws, zk.GetZooJson("/"))
}
