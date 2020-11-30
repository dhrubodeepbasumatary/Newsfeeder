package handler


import (
	 "github.com/gin-gonic/gin"
	 "net/http"
	 "rest_api_go/platform/newsfeed"

)

type newsfeedPostRequest struct{

	Title string `json:"title"`
	Post string `json:"post"`
}

func NewsfeedPost(feed *newsfeed.Repo) gin.HandlerFunc{
	return func(c *gin.Context){
		requestBody:=newsfeedPostRequest{}
		c.Bind(&requestBody)

		item:= newsfeed.Item{
			Title:requestBody.Title,
			Post:requestBody.Post,
		}
		feed.Add(item)
 		c.Status(http.StatusNoContent)
	}
}