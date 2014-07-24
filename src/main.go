package main

import (
	"code.google.com/p/go.net/websocket"
	"controllers/json"
	"controllers/wbsocket"
	"github.com/astaxie/beego"
	"log"
	"net/http"
)

func main() {
	port := "8000"
	beego.
		http.Handle("/", websocket.Handler(wbsocket.EchoMessage))
	http.HandleFunc("/getJson", json.EchoMenu)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
