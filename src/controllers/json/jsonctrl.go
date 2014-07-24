package json

import (
	//"fmt"
	"html/template"
	//"log"
	"net/http"
	//"os"
	//zk "zookeeper.fxl.com"
)

func EchoMenu(w http.ResponseWriter, r *http.Request) {
	t := template.New("menus.json")
	t.ParseFiles("views/menus.json")
	//t, _ = t.Parse("/views/menus.json")
	t.Execute(w, nil)
	//fmt.Fprintln(w, string(zk.GetZooJson("/")))
	/*fileName := "menus.json"
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
	fmt.Println(string(buf))*/
}
