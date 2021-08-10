package router

import (
	"Fusu/login"
	"Fusu/ws"
	"fmt"
	"github.com/gin-gonic/gin"
)



func Start() {
	r := gin.Default()
	r.Static("/static","./static")
	r.StaticFile("/", "./views/index.html")
	r.StaticFile("/register","./views/register.html")
	r.POST("/login", login.Login)
	r.POST("/reg", login.Register)
	r.GET("/room", ws.EchoMessage)
	err := r.Run("")
	if err != nil {
		fmt.Println("Run error...",err)
		return 
	}
}
