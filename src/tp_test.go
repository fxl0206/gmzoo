package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"testing"
)

/*func Test(t *testing.T) {
	data := []template.FuncMap{}
	data = append(data, template.FuncMap{"name": "dotcoo1", "url": "http://www.dotcoo.com/", "latlng": Latlng{24.1, 135.1}})
	data = append(data, template.FuncMap{"name": "dotcoo2", "url": "http://www.dotcoo.com/", "latlng": Latlng{24.2, 135.2}})
	data = append(data, template.FuncMap{"name": "dotcoo2", "url": "http://www.dotcoo.com/", "latlng": Latlng{24.3, 135.3}})

	datatpl := `{{range .}}{{template "user" .}}{{end}}`
	usertpl := `{{define "user"}}name:{{.name}}, url:{{.url}}, latlng:{{.latlng}} lat:{{.latlng.Lat}} lng:{{.latlng.Lng}}
{{end}}`

	tpl, err := template.New("data").Parse(datatpl)
	if err != nil {
		panic(err)
	}
	_, err = tpl.Parse(usertpl)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	println()
}*/
func Test1(t *testing.T) {
	tpl := template.New("menu.html")
	tpl.ParseFiles("views/menu.html")
	//tpl, _ = t.Parse("/views/menus.json")
	//tpl := template.New("../../views/menus.html")
	//tpl.ParseFiles("../../views/menus.html")
	//t.Execute(w, template.FuncMap{"socketServer": "ws://127.0.0.1:8000/ws"})
	log.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	log.Println("%v", tpl)
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}

type Latlng struct {
	Lat float32
	Lng float32
}

func (latlng Latlng) String() string {
	return fmt.Sprintf("%g/%g", latlng.Lat, latlng.Lng)
}
