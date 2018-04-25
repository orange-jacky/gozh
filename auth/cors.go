package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
	"fmt"
	"net/http"
)

func CORS(whiteList map[string]bool) gin.HandlerFunc  {
	// get white list and check the request url is usable
	m := make(map[string]string)
	for k, v := range whiteList {
		if v {
			m[k] = k
		}
	}
	// returns a gin handlerFunc
	return func(c *gin.Context) {

		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		// get request header's
		var headerKeys []string
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != ""{
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		}else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		// check the request is or not in white list
		if _, ok := m[origin]; ok {
			fmt.Printf("origin is not empty: %s \n", origin)
			c.Header("Access-Control-Allow-Origin", m[origin])
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			//c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		// allow all of "options" request
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}