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
func LoadRunMode()  {
	data, _ := ioutil.ReadFile("conf/yaml/RunMode.yml")
	var mode  map[string]*RunMode
	error:=yaml.Unmarshal(data, &mode)
	if error!=nil {
		fmt.Println(error)
	}
}
