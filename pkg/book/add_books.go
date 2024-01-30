package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.doublematt.librarysys/LibrarySys/models"
)

func (h handler) get_books(c *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &books)
}
