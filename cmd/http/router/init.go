package router

import (
	"webserice/initialize/configs"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Init() {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}

func Run() {
	Init()
	r.Run(configs.Setting.Port) // listen and serve on 0.0.0.0:8080
}
