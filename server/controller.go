package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Print(err)
	}
	return strings.Replace(dir, "\\", "/", -1) + "/"
}
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Ip": GetIP(),
	})
}

func GetIP() string {
	file, err := os.OpenFile("./ip.lock", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Printf("打开ip.lock失败 %s", err)
	}
	defer file.Close()

	ipLock, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("读取ip.lock失败 %s", err)
	}

	return string(ipLock)
}

func SendIp(c *gin.Context) {
	ip := c.ClientIP()
	log.Printf("send ip: %s", ip)
	clientKey := c.PostForm("clientKey")
	if clientKey != "shdgi324#@$@%@#DASFDS" {
		log.Printf("send clientKey: %s", clientKey)
		c.String(http.StatusNotAcceptable, "身份验证失败！")
		return
	}
	file, err := os.OpenFile("./ip.lock", os.O_WRONLY, os.ModePerm)

	if err != nil {
		log.Printf("打开ip.lock失败 %s", err)
		c.String(http.StatusInternalServerError, "打开ip.lock失败 %s", err)
		return
	}

	defer file.Close()

	data, err := ioutil.ReadFile("./ip.lock")

	if err != nil {
		log.Printf("读取ip.lock失败 %s", err)
		c.String(http.StatusInternalServerError, "读取ip.lock失败 %s", err)
		return
	}

	if string(data) != ip {
		_, err = file.WriteString(ip)
		if err != nil {
			log.Printf("写入ip.lock失败 %s", err)
			c.String(http.StatusInternalServerError, "写入ip.lock失败 %s", err)
			return
		}
	}
	c.String(http.StatusOK, "ok")

}
