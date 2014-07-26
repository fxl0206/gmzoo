package wbserver

import (
	//"code.google.com/p/go.net/websocket"
	"controllers/json"
	//"controllers/wbsocket"
	"log"
	"net/http"
)

type WbServer struct {
	Port        string
	MoniterPath string
}

func StaticServer(w http.ResponseWriter, req *http.Request) {
	staticHandler := http.FileServer(http.Dir("/views"))
	staticHandler.ServeHTTP(w, req)
	return
}
func (this *WbServer) Start() {
	//http.Handle("/", websocket.Handler(wbsocket.EchoMessage))
	http.HandleFunc("/getJson", json.EchoMenu)
	http.HandleFunc("/views", StaticServer)
	err := http.ListenAndServe(":"+this.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
