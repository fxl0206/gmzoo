package zook

import (
	"log"
	"testing"
)

func Test_conn(t *testing.T) {
	log.Println("ok")
	conn := GetZConn()
	log.Println(conn)
	log.Println(conn.GetJsonByPath("/"))
}
