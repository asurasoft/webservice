package http

import (
	"fmt"
	"webserice/initialize/configs"
)

func Run() {
	configs.Init()
	fmt.Println(configs.Setting.GetString("", "app_name"))
}
