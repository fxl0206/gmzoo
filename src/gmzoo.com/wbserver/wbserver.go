package wbserver

import (
	"code.google.com/p/go.net/websocket"
	"controllers/json"
	"controllers/zkvalue"
	"fmt"
	"github.com/widuu"
	"gmzoo.com/wsconn"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var ConfPath string
var socketPath string

func NewInstance(confPath string, serverName string) WbServer {
	conf := goini.SetConfig(confPath)
	port := conf.GetValue(serverName, "port")
	viewPath := conf.GetValue(serverName, "viewPath")
	ConfPath = confPath
	wsPath := conf.GetValue("websocket", "wsPath")
	wsPort := conf.GetValue("websocket", "port")
	wsIp := conf.GetValue("websocket", "wsIp")
	socketPath = "ws://" + wsIp + ":" + wsPort + "/" + wsPath
	return WbServer{port, viewPath, serverName}
}

type WbServer struct {
	port       string
	viewPath   string
	ServerName string
}

func (this *WbServer) Start() {
	http.Handle("/ws", websocket.Handler(DealSocketReq))
	http.HandleFunc("/"+this.viewPath+"/", StaticServer)
	http.HandleFunc("/getMenu", json.EchoMenu)
	http.HandleFunc(zkvalue.Prefix+"/", zkvalue.EchoValue)
	err := http.ListenAndServe(":"+this.port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
func StaticServer(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/"):]
	log.Println(r.URL.Path)
	if m, _ := regexp.MatchString("/json", r.URL.Path); m {
		w.Header().Add("Content-Type", "application/json")
		log.Println(w.Header().Get("Content-Type"))
	}
	if m, _ := regexp.MatchString("/menu.html", filePath); m {
		t := template.New("menu.html")
		t.ParseFiles(filePath)
		t.Execute(w, socketPath)
	} else {
		http.ServeFile(w, r, filePath)
	}
}
func DealSocketReq(ws *websocket.Conn) {
	wsconn.PutConn(ws)
	var err error
	websocket.Message.Send(ws, "websocket Connection successfull!")
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
