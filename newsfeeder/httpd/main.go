package main

import (
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	Port int = 8080
)

func main() {
	feed := newsfeed.New()

	r := gin.Default() //creates new Gin, combined with logger, and recovery middleware which catches endpoints if it panics
	
	r.GET("/nice ports,/Trinity.txt.bak", handler.ChangePort())

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	portConfig := ":" + strconv.Itoa(Port)
	r.Run(portConfig) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
