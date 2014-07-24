package main

import (
	"code.google.com/p/go.net/websocket"
	zk "com/fxl/zookeeper"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getJson(w http.ResponseWriter, r *http.Request) {
	fileName := "menus.json"
	fl, err := os.Open(fileName)
	defer fl.Close()
	if err != nil {
		log.Fatal("www", err)
	}
	fmt.Fprintln(w, string(zk.GetZooJson("/")))
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
	}
	fmt.Println(string(buf))
}
func onConnect(ws *websocket.Conn) {
	websocket.JSON.Send(ws, zk.GetZooJson("/"))
}
func main() {
	http.Handle("/chat", websocket.Handler(onConnect))
	http.HandleFunc("/getJson", getJson)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
