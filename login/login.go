package login

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func init() {

}

type userLogin struct {
	Account string `form:"account" json:"account"`
	Password  string `form:"password" json:"password"`
}

type userRegister struct {
	Account string `form:"account" json:"account"`
	Nick string `form:"nick" json:"nick"`
	Password  string `form:"password" json:"password"`
}

var DB, _ = sql.Open("mysql","root:324243@tcp(localhost:3305)/Fusu")

func Register(c *gin.Context) {
	var form userRegister
	err:=c.Bind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bind error": err.Error()})
		return
	}

	tx, err := DB.Begin()
	stmt, err := tx.Prepare("insert user set account=?,password=?,nick=?")
	if err != nil {
		println("prepare error",err)
		return
	}
	//fmt.Println(form.Account,form.Password,form.Nick)
	_, err = stmt.Exec(form.Account,form.Password,form.Nick)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Exec error": err.Error()})
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}

	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/")
}

func Login(c *gin.Context) {
	var form userLogin
	var dist userRegister
	err := c.Bind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"bind error": err.Error()})
		return
	}
	tx, _ :=DB.Begin()

	err = tx.QueryRow("select * from user where account = ?",form.Account).Scan(&dist.Account,&dist.Password,&dist.Nick)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"用户不存在": err.Error()})
		return
	}
	http.ServeFile(c.Writer,c.Request,"views/room.html")

	if err != nil {
		return
	}
}
