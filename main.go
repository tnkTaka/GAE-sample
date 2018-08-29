package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	ApiRouter(r.Group("/api"))
	r.Run(":8080")
}

func ApiRouter(api *gin.RouterGroup)  {
	api.GET("/",GetID)
}

func GetID(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"text":"Hello World"})
}