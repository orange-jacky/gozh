package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Demo(c *gin.Context) {
	view := "view/demo.html"
	data := gin.H{"user": c.ClientIP()}
	c.HTML(http.StatusOK, view, data)
}
