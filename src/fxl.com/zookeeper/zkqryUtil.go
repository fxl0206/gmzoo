package zookeeper

import (
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
func qryNode(id int, rPath string, c *zk.Conn) znode {

	children, _, _, err := c.ChildrenW(rPath)
	ret := znode{id, rPath, rPath, []znode{}}

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
func GetZooJson(path string) []byte {
	Zc, _, err := zk.Connect([]string{"115.29.8.106"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
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
