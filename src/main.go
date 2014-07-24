package main

import (
	"code.google.com/p/go.net/websocket"
	"controllers/json"
	"controllers/wbsocket"
	"log"
	"net/http"
)

func main() {
	port := "8000"
	http.Handle("/", websocket.Handler(wbsocket.EchoMessage))
	http.HandleFunc("/getJson", json.EchoMenu)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
