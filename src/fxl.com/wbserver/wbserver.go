package wbserver

import (
	"code.google.com/p/go.net/websocket"
	"controllers/json"
	//"controllers/wbsocket"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type WbServer struct {
	Port        string
	MoniterPath string
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if m, _ := regexp.MatchString("/json", r.URL.Path); m {
		w.Header().Add("Content-Type", "application/json")
		log.Println(w.Header().Get("Content-Type"))
	}
	http.ServeFile(w, r, r.URL.Path[len("/"):])
}
func (this *WbServer) Start() {
	http.Handle("/ws", websocket.Handler(ChatWith))
	http.HandleFunc("/views/", StaticServer)
	http.HandleFunc("/getMenu", json.EchoMenu)
	err := http.ListenAndServe(":"+this.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
func ChatWith(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		//msg := "Received from " + ws.Request().Host + "  " + reply
		msg := "welcome to websocket do by pp"
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
