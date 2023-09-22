package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Run()
}

func checkConnection(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Pong"})
}
