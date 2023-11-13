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
	router.GET("/books", getAllBooks)
	router.DELETE("/book/:id", deleteBookById)
	router.POST("/book", addBook)

	router.Run()
}

func checkConnection(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Pong"})
}

// temporary db for development purpose

var books = []Book{
	{Id: 1, Title: "The beginning", Author: "Dan Brown"},
	{Id: 2, Title: "Harry Potter and philosopher stone", Author: " J. K. Rowling"},
	{Id: 3, Title: "Sword of Destiny", Author: "Jacek Sapkowski"},
	{Id: 4, Title: "The Blood of Strangers", Author: "Frank Hyuler"},
}

//
// CRUD methods
//

func getAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookBtId(c *gin.Context) {

	// converting string param to int value
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, findBook(id))

}

func deleteBookById(c *gin.Context) {

	// converting string param to int value
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	if deleteBook(id) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "Deleted"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status": "Book not found"})
	}
}

func addBook(c *gin.Context) {

	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		panic("error while binding JSON to book object")
	}

	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)

}

//
// secondary methods
//

func updateBookById(id int, updatedBook Book) {

	if findBook(id).Id != 0 {
		for i, book := range books {
			if book.Id == id {
				var books_copy []Book
				books_copy = append(books[:i], updatedBook)
				books = append(books_copy, books[i+1:]...)
				break
			}
		}
	}

}

func findBook(id int) Book {

	for _, book := range books {
		if book.Id == id {
			return book
		}

	}
	return Book{Id: 0, Title: "Error", Author: "Book not Found!", Pages: 0}
}

func deleteBook(id int) bool {

	for i, book := range books {
		if book.Id == id {
			books = append(books[:i], books[i+1:]...)
			return true
		}
	}

	return false
}
