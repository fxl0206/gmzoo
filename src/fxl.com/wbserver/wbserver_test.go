package wbserver

import (
	"log"
	"net/http"
)

func Test_showdir() {
	log.Println(http.Dir("/views"))
}
