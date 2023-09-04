package routes

import (
	"canonicalTestAPI/pkg/handlers"
	"canonicalTestAPI/pkg/service"

	"github.com/gin-gonic/gin"
)

// configures the server's endpoints and database middleware
func Router(serv service.Service) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("service", serv)
		c.Next()
	})

	router.POST("/book", handlers.CreateBook)
	router.GET("/books", handlers.Query)
	router.GET("/book/:id", handlers.GetBook)
	router.DELETE("/book/:id", handlers.DeleteBook)
	router.PATCH("/book/:id", handlers.UpdateBook)
	router.POST("/collection", handlers.CreateCollection)
	router.PATCH("/collection/:name", handlers.UpdateCollection)
	router.GET("/collections", handlers.GetAllCollections)
	router.DELETE("/collection/:name", handlers.DeleteCollection)
	router.POST("/collection/:name/:id", handlers.CollectBook)
	router.DELETE("/collection/:name/:id", handlers.DiscardBook)

	return router
}
