package wsconn

import (
	"code.google.com/p/go.net/websocket"
	"sync"
)

var WsConnSlise = []*websocket.Conn{}

type Wsconn struct {
	onlineCt int64
	lock     sync.Mutex
}

func (ws *Wsconn) SendMsg(msg string) {
	for _, conn := range WsConnSlise {
		websocket.Message.Send(conn, msg)
	}
}
func PutConn(conn *websocket.Conn) {
	WsConnSlise = append(WsConnSlise, conn)
}
