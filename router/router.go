package router

import (
	"github.com/gin-gonic/gin"
)



func Start() {
	r := gin.Default()
	r.Static("/static","./static")
	r.StaticFile("/", "./views/index.html")
	//fmt.Println("init...")
	r.StaticFile("/room", "./views/room.html")

	r.POST("/login", func(context *gin.Context) {

	})
	err := r.Run()
	if err != nil {
		return 
	}
}
