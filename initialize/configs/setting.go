package configs

import (
	"fmt"
	"sync"
	"webserice/utils"
)

var Setting utils.IniParser

var once sync.Once

func Init() {
	once.Do(func() {
		conf_file_name := "config/config.yml"
		if err := Setting.Load(conf_file_name); err != nil {
			fmt.Printf("try load config file[%s] error[%s]\n", conf_file_name, err.Error())
			return
		}
	})
}
