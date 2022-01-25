package main

import (
	"What2Buy/Server/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.PingHandler)

	r.GET("/getItems", handler.GetItems)
	r.POST("/addItem", handler.AddItem)
	r.POST("/upvoteItem", handler.UpvoteItem)
	r.DELETE("/deleteItem", handler.DeleteItem)

	r.Run() // listen and serve on localhost:8080
}
