package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
func LoadProjectConf()  {
	LoadRunMode()
}

type RunMode struct {
	Addr string
	Port string
}
//
var Mode  map[string]*RunMode
func LoadRunMode()  {
	data, _ := ioutil.ReadFile("conf/yaml/RunMode.yml")
	error:=yaml.Unmarshal(data, &Mode)
	if error!=nil {
		fmt.Println(error)
	}
}
