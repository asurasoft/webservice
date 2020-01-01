package configs

import "fmt"

type AppSetting struct {
	Port    string
	AppName string
}

var Setting = new(AppSetting)

func Init() {
	InitParse()

	Setting.Port = parse.GetString("", "port")
	Setting.AppName = parse.GetString("", "app_name")
	fmt.Println(Setting)

}
