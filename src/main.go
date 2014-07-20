package main

import (
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
	//fmt.Fprintln(w, string(buf))

}
func main() {
	http.HandleFunc("/getJson", getJson)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
