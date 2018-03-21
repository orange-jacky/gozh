package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Testpost(c *gin.Context) {
	view := "view/testpost.html" //选用哪个视图模板

	type D struct {
		Ip, User, Passwd string
	}
	d := &D{}
	d.Ip = c.ClientIP()
	d.User = c.PostForm("username")
	d.Passwd = c.PostForm("password")

	c.HTML(http.StatusOK, view, d)
}
