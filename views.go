package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func listTvShows(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tvShows)
}

func addTvShow(c *gin.Context) {
	var tvShow series

	if err := c.BindJSON(&tvShow); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	tvShows = append(tvShows, tvShow)
	c.IndentedJSON(http.StatusCreated, tvShow)
}

func getTvShowbyID(id string) (int, *series) {
	int_id, err := strconv.Atoi(id) // Convert id from string to integer
	if err != nil {
		panic(err)
	}

	for index, tvShow := range tvShows {
		if tvShow.ID == int_id {
			return index, &tvShow
		}
	}

	return -1, nil
}

func getTvShow(c *gin.Context) {
	id := c.Param("id")
	_, tvShow := getTvShowbyID(id)
	if tvShow == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TV Show not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, tvShow)
}

func deleteTvShow(c *gin.Context) {
	id := c.Param("id")
	index, _ := getTvShowbyID(id)
	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TV Show not found"})
		return
	}
	tvShows = append(tvShows[:index], tvShows[index+1:]...)
	c.Status(http.StatusNoContent)
}
