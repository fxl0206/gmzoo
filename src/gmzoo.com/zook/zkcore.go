package zook

import (
	"conf"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/widuu"
	"gmzoo.com/node"
	"gmzoo.com/wsconn"
	"log"
	"sync"
	"time"
)

var zconn ZkCore

func GetZConn() ZkCore {
	if zconn.isInit {
		return zconn
	} else {
		zconn = ZkCore{evtCache: make(map[string]<-chan zk.Event), isInit: true}
		zconn.initial()
	}
	return zconn
}

type ZkCore struct {
	conn      *zk.Conn
	socktConn *wsconn.Wsconn
	cacheLock sync.Mutex
	evtCache  map[string]<-chan zk.Event
	isInit    bool
}

func (k *ZkCore) initial() {
	k.conn = k.getNewConnection()
	k.cacheLock = sync.Mutex{}

}
func (k *ZkCore) GetJsonByPath(path string) string {
	node := k.qryNode(0, path)
	return node.ToJsonStr()
}
func (k *ZkCore) nodeChage(zh <-chan zk.Event) {
	e := <-zh
	wType := e.Type
	log.Println("%V", wType)
	path := e.Path
	k.clearEvtCache(path)
	switch wType {
	case zk.EventNodeChildrenChanged:
		jData := k.GetJsonByPath(path)
		k.socktConn.SendMsg(jData)
	case zk.EventNodeDataChanged:
		k.socktConn.SendMsg("data change:" + path)
	case zk.EventNodeDeleted:
		k.socktConn.SendMsg("info: node [" + path + "] is deleted!")
	case zk.EventNodeCreated:
		k.socktConn.SendMsg("node created:" + path)
	default:
		k.socktConn.SendMsg("node default:" + path)
	}

}
func (k *ZkCore) setEvtCache(path string, zEvt <-chan zk.Event) {
	k.cacheLock.Lock()
	defer k.cacheLock.Unlock()
	k.evtCache[path] = zEvt
}
func (k *ZkCore) clearEvtCache(path string) {
	k.setEvtCache(path, nil)
}
func (k *ZkCore) containCacheKey(path string) bool {
	chkRes := false
	k.cacheLock.Lock()
	defer k.cacheLock.Unlock()
	if k.evtCache[path] != nil {
		chkRes = true
	}
	return chkRes

}
func (k *ZkCore) qryNode(id int, rPath string) node.Znode {
	c := k.GetConnection()
	children, _, nodeEvt, err := c.ChildrenW(rPath)
	ret := node.Znode{id, rPath, rPath, []node.Znode{}}
	if err != nil {
		log.Println(err)
		k.setEvtCache(rPath, nil)
	} else {
		if !k.containCacheKey(rPath) {
			go k.nodeChage(nodeEvt)
		}
		k.setEvtCache(rPath, nodeEvt)
		for i, child := range children {
			fixPath := rPath
			if fixPath == "/" {
				fixPath = ""
			}
			cPath := fixPath + "/" + child
			cRet := k.qryNode(i, cPath)
			ret.AddChild(cRet)
		}
	}
	return ret
}

func (k *ZkCore) GetConnection() *zk.Conn {
	if k.conn != nil {
		return k.conn
	} else {
		stat := k.conn.State()
		log.Println(stat)
		switch stat {
		case zk.StateUnknown:
			log.Print("new connection:")
			k.conn = k.getNewConnection()
			log.Println(stat)
		}
	}
	return k.conn
}
func (k *ZkCore) getNewConnection() *zk.Conn {
	serverName := "zookeeper"
	cf := goini.SetConfig(conf.CurConfFile)
	ip := cf.GetValue(serverName, "ip")
	port := cf.GetValue(serverName, "port")
	c, _, err := zk.Connect([]string{ip + ":" + port}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	return c
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
