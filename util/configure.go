package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

//Gin
type Gin struct {
	Mode            string `json:"mode"`
	Url             string `json:"url"`
	Port            int    `json:"port"`
	Timeout_read_s  int    `json:"timeout_read_s"`
	Timeout_write_s int    `json:"timeout_write_s"`
}

// Log 保存日志配置信息
type Log struct {
	File   string `json:"file"`
	Access string `json:"access"`
}

//es
type Es struct {
	Hosts string `json:"hosts"`
}

//mongo
type Mongo struct {
	Hosts             string `json:"hosts"`
	Connect_timeout_s int    `json:"connect_timeout_s"`
	Username          string `json:"username"`
	Passwd            string `json:"passwd"`
	DatabaseName      string `json:"database_name"`
}

//pic_addr
type Pic_addr struct {
	Prefix string `json:"prefix"`
}

//collections' names
type Collections_names struct {
	User     string `json:"user"`
	Ariticle string `json:"ariticle"`
}


//configure
type configure struct {
	Gin               Gin               `json:"gin"`
	Log               Log               `json:"log"`
	Es                Es                `json:"es"`
	Mongo             Mongo             `json:"mongo"`
	Pic_addr          Pic_addr          `json:"pic_addr"`
	Collections_names Collections_names `json:"collections_names"`
	WhiteList	map[string]bool 		`json:"white_list"`
}

var (
	conf      *configure
	conf_once sync.Once
)

//Configure 载入json配置文件
func Configure(file string) *configure {
	conf_once.Do(func() {
		conf = &configure{}
		if err := conf.init(file); err != nil {
			log.Fatalln(err)
		}
	})
	return conf
}

//init 载入json配置文件
func (c *configure) init(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("error open file %s fail,%v", file, err)
	}
	defer fd.Close()
	/*
		//处理不了字段中包含html字符的,比如&
		content, err := ioutil.ReadAll(fd)
		if err != nil {
			return err
		}
		return xml.Unmarshal(content, c)
	*/
	//使用decoder处理包含html字符的内容

	//d := xml.NewDecoder(fd)
	//d.Strict = false
	//d.AutoClose = xml.HTMLAutoClose
	//d.Entity = xml.HTMLEntity

	b := make([]byte, 1000)
	n, err := fd.Read(b)
	if n == 0 {
		fmt.Print("配置文件中未找到相应数据")
	}
	return json.Unmarshal(b[:n], c)
}

func (c *configure) String() string {
	js, _ := json.MarshalIndent(c, "", "\t")
	return fmt.Sprintf("%s", js)
}

//得到配置实例
func GetConfigure() *configure {
	return conf
}
