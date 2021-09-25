package main

import (
		"net/http"
		"github.com/gin-gonic/gin"
)

func main() {
		router := gin.Default()

		router.GET("/", rootHandler)

		router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "test",
		"bio": "engineer",
	})
}