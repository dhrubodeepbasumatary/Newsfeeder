package main

import (
	"github.com/gin-gonic/gin"
	"rest_api_go/httpd/handler"
	"rest_api_go/platform/newsfeed"
	//"fmt"
)
func main() {
	feed:=newsfeed.New()
	r := gin.Default()
	r.GET("/ping",handler.PingGet())
	r.GET("/newsfeed",handler.NewsfeedGet(feed))
	r.POST("/newsfeed",handler.NewsfeedPost(feed))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}