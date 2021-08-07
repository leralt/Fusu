package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func Start() {
	r := gin.Default()
	r.Static("/static","./static")
	r.StaticFile("/", "./views/index.html")
	//fmt.Println("init...")
	r.StaticFile("/room", "./views/room.html")
	r.StaticFile("/register","./views/register.html")
	r.POST("/login", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		if username == "long" && password == "123" {
			context.Redirect(http.StatusMovedPermanently, "http://localhost:8080/room")
		}
	})
	err := r.Run()
	if err != nil {
		return 
	}
}
