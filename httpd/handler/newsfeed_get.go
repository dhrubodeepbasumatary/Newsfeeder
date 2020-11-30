package handler

import (
	 "github.com/gin-gonic/gin"
	 "rest_api_go/platform/newsfeed"
	 "net/http"
)
func NewsfeedGet(feed *newsfeed.Repo) gin.HandlerFunc{
	return func(c *gin.Context){
		results:=feed.GetAll()
 		c.JSON(http.StatusOK,results)
	}
}