package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var server string
var port string
var clientKey string

func main() {
	getConfig()
	for t := range time.Tick(time.Second * 5) {
		fmt.Printf("%s sendIp \n", t.Local())
		resp, err := http.PostForm("http://"+server+":"+string(port)+"/sendIp",
			url.Values{"clientKey": {clientKey}})
		if err != nil {
			log.Printf("sendIp Error: %s", err)
			continue
		}
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("读取结果 Error: %s", err)
			continue
		}
		if string(content) == "ok" {
			log.Printf("send ok")
		}
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Print(err)
	}
	return strings.Replace(dir, "\\", "/", -1) + "/"
}

func getConfig() {
	confPath := getCurrentDirectory() + "conf"
	log.Printf("config path : %s", confPath)
	file, err := os.Open(confPath)
	if err != nil {
		log.Printf("打开配置文件conf失败!")
		return
	}

	defer file.Close()

	conf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("读取配置文件失败!")
		return
	}

	confStr := string(conf)
	confs := strings.Split(confStr, "\n")
	for i := range confs {
		items := strings.Split(string(confs[i]), "=")
		if len(items) < 2 {
			log.Printf("配置文件不符合规范!")
		}
		key := strings.Replace(items[0], " ", "", -1)
		value := strings.Replace(items[1], " ", "", -1)
		if key == "server" {
			server = value
		}

		if key == "port" {
			port = value
		}

		if key == "clientKey" {
			clientKey = value
		}
	}
}
