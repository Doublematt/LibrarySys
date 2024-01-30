package book

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func register_routes(router *gin.Engine, db *gorm.DB) {

	h := &handler{
		DB: db,
	}

	routes := router.Group("/books")
	routes.GET("/", h.get_books)

}
