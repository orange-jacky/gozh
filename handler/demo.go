package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Demo(c *gin.Context) {

	data := gin.H{"user": c.ClientIP()}
	c.String(http.StatusOK, "%v", data)
}
