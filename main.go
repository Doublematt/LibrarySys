package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.doublematt.librarysys/LibrarySys/models"
)

// init models
type Book models.Book

func main() {
	router := gin.Default()

	// all endpoints available
	// check connection
	router.GET("/ping", checkConnection)
	router.GET("/book/:id", getBookBtId)

	router.Run()
}

//
// CRUD methods
//

func checkConnection(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Pong"})
}

func getBookBtId(c *gin.Context) {

	// converting string param to int value
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, findBook(id))

}

// temporary db for development purpose

var books = []Book{
	{Id: 1, Title: "The beginning", Author: "Dan Brown"},
}

//
// secondary methods
//

func findBook(id int) Book {

	for _, book := range books {
		if book.Id == id {
			return book
		}

	}
	return Book{Id: 0, Title: "Error", Author: "Book not Found!", Pages: 0}
}
