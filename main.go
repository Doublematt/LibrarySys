package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.doublematt.librarysys/LibrarySys/models"
)

type Book models.Book

func main() {
	router := gin.Default()

	// check connection
	router.GET("/ping", checkConnection)

	router.Run()
}

func checkConnection(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Pong"})
}

var books = []Book{
	{Id: 1, Title: "The beginning", Author: "Dan Brown"},
}
