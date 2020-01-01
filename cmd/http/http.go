package http

import (
	"webserice/cmd/http/router"
	"webserice/initialize/configs"
)

func Run() {
	configs.Init()
	router.Run()
}
