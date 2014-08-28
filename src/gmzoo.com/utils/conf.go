package utils

import (
	"github.com/widuu"
	"log"
)

func GetConf(filePath string) Conf {
	conf := Conf{filePath}
	conf.init()
	return conf
}

type Conf struct {
	filePath string
}

var confSections []map[string]map[string]string

func (this *Conf) init() {
	cf := goini.SetConfig(this.filePath)
	confSections = cf.ReadList()
}
func (this *Conf) GetWebPort() string {
	return this.GetValue("webserver", "port")
}
func (this *Conf) GetWebViewPath() string {
	return this.GetValue("webserver", "viewPath")
}
func (this *Conf) GetValue(section, name string) string {
	for _, v := range confSections {
		for key, value := range v {
			if key == section {
				return value[name]
			}
		}
	}
	panic("config item" + name + " is not exists!")
	return "no value"
}
