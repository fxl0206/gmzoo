package zkvalue

import (
	"fmt"
	zk "fxl.com/utils"
	"log"
	"net/http"
)

var Prefix string = "/getValue"

func EchoValue(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len(Prefix):]
	log.Println(path)
	conn := zk.GetConnection()
	value, _, err := conn.Get(path)
	if err != nil {
		fmt.Fprintln(w, err)
		log.Println(err)
	} else {
		w.Write(value)
	}
}
