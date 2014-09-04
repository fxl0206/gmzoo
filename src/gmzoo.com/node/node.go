package node

import (
	"encoding/json"
	"log"
)

type Znode struct {
	Id     int
	Name   string
	Url    string
	Childs []Znode
}

func (n *Znode) AddChild(node Znode) {
	n.Childs = append(n.Childs, node)
}
func (n *Znode) ToJsonStr() string {
	v, err := json.Marshal(n)
	if err != nil {
		log.Println(err)
	}
	return string(v)
}
