package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.LoadHTMLGlob(getCurrentDirectory() + "templates/*")
	v1 := router.Group("/")
	{
		v1.GET("/index.html", IndexPage)
		v1.POST("/sendIp", SendIp)
	}

	router.Run(":9010")
}
