package zkUtil

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/widuu"
	"log"
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
func (n *znode) ToJsonStr() string {
	v, err := json.Marshal(n)
	if err != nil {
		log.Println(err)
	}
	return string(v)
}

var WsConnSlise = []*websocket.Conn{}
var WsSlise *websocket.Conn

func nodeChage(zh <-chan zk.Event) {
	e := <-zh
	wType := e.Type
	log.Println("%V", wType)
	path := e.Path
	EvtCache[path] = nil
	if WsSlise != nil {
		switch wType {
		case zk.EventNodeChildrenChanged:
			jData := GetZooJson(path)
			websocket.Message.Send(WsSlise, jData)
		case zk.EventNodeDataChanged:
			websocket.Message.Send(WsSlise, "data change:"+path)
		case zk.EventNodeDeleted:
			websocket.Message.Send(WsSlise, "info: node ["+path+"] is deleted!")
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
	if err != nil {
		log.Println(err)
		EvtCache[rPath] = nil
	} else {
		if EvtCache[rPath] == nil {
			go nodeChage(nodeEvt)
		}
		EvtCache[rPath] = nodeEvt
		for i, child := range children {
			fixPath := rPath
			if fixPath == "/" {
				fixPath = ""
			}
			cPath := fixPath + "/" + child
			cRet := qryNode(i, cPath, c)
			ret.addChild(cRet)
		}
	}
	return ret
}

var Zc *zk.Conn = nil

func GetConnection() *zk.Conn {
	if Zc == nil {
		Zc = getNewConnection()
	} else {
		stat := Zc.State()
		log.Println(stat)
		switch stat {
		case zk.StateUnknown:
			log.Print("new connection:")
			Zc = getNewConnection()
			log.Println(stat)
		}
	}
	return Zc
}
func getNewConnection() *zk.Conn {
	serverName := "zookeeper"
	cf := goini.SetConfig("")
	ip := cf.GetValue(serverName, "ip")
	port := cf.GetValue(serverName, "port")
	c, _, err := zk.Connect([]string{ip + ":" + port}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	return c
}
func GetZooJson(path string) string {
	Zc = GetConnection()
	node := qryNode(0, path, Zc)
	return node.ToJsonStr()
}

/*
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}
	prePath := SubString(path, 0, strings.LastIndex(path, "/"))
	if prePath == "" {
		prePath = "/"
	}
	jData := GetZooJson(prePath)
	fmt.Println("!!  " + prePath + "  !!!" + jData)
*/
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
