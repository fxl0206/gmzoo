package zkUtil

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"os"
	"time"
)

type znode struct {
	Id     int
	Name   string
	Url    string
	Childs []znode
}

func (n *znode) addChild(node znode) {
	n.Childs = append(n.Childs, node)
}

var WsConnSlise = []*websocket.Conn{}
var WsSlise *websocket.Conn

func nodeChage(zh <-chan zk.Event) {
	e := <-zh
	wType := e.Type
	fmt.Println("%V", wType)
	path := e.Path
	if WsSlise != nil {
		fmt.Println("%V", path)
		switch wType {
		case zk.EventNodeChildrenChanged:
			Zc = getConnection()
			children, _, nodeEvt, err := Zc.ChildrenW(path)
			if err != nil {
				log.Println(err)
				EvtCache[path] = nil
			} else {
				EvtCache[path] = nodeEvt
				go nodeChage(nodeEvt)
				for _, child := range children {
					if path == "/" {
						path = ""
					}
					cPath := path + "/" + child
					if EvtCache[cPath] == nil {
						_, _, evt, err := Zc.ChildrenW(cPath)
						if err != nil {
							log.Println(err)
						}
						EvtCache[cPath] = evt
						websocket.Message.Send(WsSlise, "{\"type\":\"add\",\"path\":\""+cPath+"\"}")
						go nodeChage(evt)
					}
				}
			}
		case zk.EventNodeDataChanged:
			websocket.Message.Send(WsSlise, "data change:"+path)
		case zk.EventNodeDeleted:
			websocket.Message.Send(WsSlise, "node deleted:"+path)
			EvtCache[path] = nil
		case zk.EventNodeCreated:
			websocket.Message.Send(WsSlise, "node created:"+path)
		default:
			websocket.Message.Send(WsSlise, "node default:"+path)
		}

	}
}

var EvtCache = make(map[string]<-chan zk.Event)

func qryNode(id int, rPath string, c *zk.Conn) znode {

	children, _, nodeEvt, err := c.ChildrenW(rPath)
	ret := znode{id, rPath, rPath, []znode{}}
	EvtCache[rPath] = nodeEvt
	go nodeChage(nodeEvt)
	if err != nil {
		log.Println(err)
	}
	for i, child := range children {
		if rPath == "/" {
			rPath = ""
		}
		cPath := rPath + "/" + child
		cRet := qryNode(i, cPath, c)
		ret.addChild(cRet)
	}
	return ret
}

var Zc *zk.Conn

func getConnection() *zk.Conn {
	//if Zc == nil {
	c, _, err := zk.Connect([]string{"115.29.8.106"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	Zc = c
	//}
	return Zc
}
func GetZooJson(path string) []byte {
	Zc = getConnection()
	group := qryNode(0, path, Zc)
	v, err := json.Marshal(group)
	if err != nil {
		fmt.Println(err)
	}
	return v
}
func main() {

	os.Stdout.Write(GetZooJson("/"))
	log.Println()

}

/*
	/*group := znode{
	1,
	"test",
	"/qq",
	[]znode{znode{2, "tt", "ee", []znode{}}}}
type ColorGroup struct {
	ID     string
	Name   string
	Colors []string
}

func main() {
	group := ColorGroup{
		ID:     "1",
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}*/
