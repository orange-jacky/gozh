package util

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

//结构要直接转换成json字符串,结构中变量首字母必须大写
type LogFormat struct {
	Time       string `json:"time"`
	Clientip   string `json:"clientip,omitempty"`
	Remote_ip  string `json:"remote_ip"`
	Method     string `json:"method"`
	Uri        string `json:"uri"`
	Query      string `json:"query"`
	Proto      string `json:"proto"`
	User_agent string `json:"user_agent"`
	Referer    string `json:"referer"`
	Host       string `json:"host"`
	Status     int    `json:"status"`
	Latency    int64  `json:"latency"`
	Bytes_in   int64  `json:"bytes_in"`
	Bytes_out  int64  `json:"bytes_out"`
}

func (format *LogFormat) Marshal() (string, error) {
	//result, err := json.Marshal(format)
	//return string(result), err
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(format); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Log format which can be constructed using the following tags:
//
// - time_rfc3339
// - remote_ip
// - uri
// - host
// - method
// - path
// - referer
// - user_agent
// - status
// - latency (In microseconds)
// - latency_human (Human readable)
// - bytes_in (Bytes received)
// - bytes_out (Bytes sent)
//
// Example "${remote_ip} ${status}"

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()

		proto := c.Request.Proto
		host := c.Request.Host
		method := c.Request.Method
		uri := c.Request.URL.Path
		remote_ip := c.Request.RemoteAddr
		sli := strings.Split(remote_ip, ":")
		remote_ip = sli[0]

		usserip := c.Request.Header.Get("X-Forwarded-For")
		if usserip == "" {
			usserip = c.Request.Header.Get("X-Real-Ip")
		}

		bytes_in := c.Request.ContentLength
		//path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		user_agent := ""
		if val, ok := c.Request.Header["User-Agent"]; ok {
			user_agent = strings.Join(val, ",")
		}
		referer := ""
		if val, ok := c.Request.Header["Referer"]; ok {
			referer = strings.Join(val, ",")
		}
		//before request
		c.Next()
		//after request
		latency := time.Now().Sub(begin).Nanoseconds() / 1000000 //unit:ms
		//latency_human := time.Since(begin).String()
		//access the status we are sending
		status := c.Writer.Status()
		bytes_out := int64(c.Writer.Size())
		//bytes_out := c.Request.Response.ContentLength

		format := &LogFormat{}
		format.Time = begin.Format("2006-01-02 15:04:05")
		format.Remote_ip = remote_ip
		format.Clientip = usserip
		format.Method = method
		format.Uri = uri
		format.Query = query
		format.Proto = proto
		format.User_agent = user_agent
		format.Referer = referer
		format.Host = host
		format.Status = status
		format.Latency = latency
		format.Bytes_in = bytes_in
		format.Bytes_out = bytes_out

		if result, err := format.Marshal(); err == nil {
			//fmt.Printf(result)
			mylog := GetMylog().Access()
			mylog.Info(result)
		}
	}
}
