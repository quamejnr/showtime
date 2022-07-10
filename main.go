package main

import (
	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	
	router.GET("/series", listTvShows)
	router.POST("/series", addTvShow)
	router.GET("/series/:id", getTvShow)
	router.DELETE("/series/:id", deleteTvShow)

	router.Run("localhost:8080")

}



