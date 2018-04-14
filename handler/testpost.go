package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Testpost(c *gin.Context) {
	type D struct {
		Ip, User, Passwd string
	}
	d := &D{}
	d.Ip = c.ClientIP()
	d.User = c.PostForm("username")
	d.Passwd = c.PostForm("password")

	c.String(http.StatusOK, "%v", d)
}
