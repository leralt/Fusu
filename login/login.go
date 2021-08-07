package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type req struct {
	Name string `json:"username"`
	Pass  string `json:"password"`
}
func Auth(c *gin.Context) {
	form := req{}
	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if form.Name == "long" && form.Pass == "123"{

	}
}