package configs

import (
	"fmt"
	"sync"
	"webserice/utils"
)

var parse utils.IniParser

var once sync.Once

func InitParse() {
	once.Do(func() {
		conf_file_name := "config/config.yml"
		if err := parse.Load(conf_file_name); err != nil {
			fmt.Printf("try load config file[%s] error[%s]\n", conf_file_name, err.Error())
			return
		}
	})
}
