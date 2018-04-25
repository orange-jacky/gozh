package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/orange-jacky/gozh/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
	"github.com/orange-jacky/gozh/auth"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

type Stoper interface {
	Stop()
}

var all []Stoper

func Usage(program string) {
	fmt.Printf("\nusage: %s conf/cf.json\n", program)
	fmt.Printf("\nconf/cf.json      configure file\n")
}

func main() {
	if len(os.Args) != 2 {
		Usage(os.Args[0])
		os.Exit(-1)
	}
	log.Println("[Main] Starting program")
	defer log.Println("[Main] Exit program successful.")

	Init(os.Args[1])
	defer Release()

	//配置gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(util.Logger())
	//router.Static("/static", "./static") //载入js,css,img等静态资源
	//router.LoadHTMLGlob("view/*.html")   //载入html模板
	conf := util.GetConfigure()
	// use CORS middleware
	router.Use(auth.CORS(conf.WhiteList))
	g := conf.Gin
	server := fmt.Sprintf(":%v", g.Port)

	prefix := fmt.Sprintf("%s", g.Url)
	AllRouter(prefix, router)

	//起一个http服务器
	s := &http.Server{
		Addr:         server,
		Handler:      router,
		ReadTimeout:  time.Duration(g.Timeout_read_s) * time.Second,
		WriteTimeout: time.Duration(g.Timeout_write_s) * time.Second,
	}
	go func(s *http.Server) {
		log.Printf("[Main] http server start\n")
		err := s.ListenAndServe()
		log.Printf("[Main] http server stop (%+v)\n", err)
	}(s)
	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	for {
		select {
		case sig := <-signals:
			log.Println("[Main] Catch signal", sig)
			//平滑关闭server
			err := s.Shutdown(context.Background())
			log.Printf("[Main] start gracefully shuts down http serve %+v", err)
			return
		}
	}
}

func SetupCPU() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
}

func Init(filename string) {
	//配置使用cpu数量
	SetupCPU()
	//读取配置文件信息
	util.Configure(filename)
	//1
	mylog := util.Mylog()
	all = append(all, mylog)
}

func Release() {
	for _, v := range all {
		v.Stop()
	}
}
