package routes

import (
	"canonicalTestAPI/pkg/handlers"
	"canonicalTestAPI/pkg/service"

	"github.com/gin-gonic/gin"
)

func Router(serv service.Service) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("service", serv)
		c.Next()
	})

	router.POST("/book", handlers.CreateBook)
	router.GET("/books", handlers.Query)
	router.DELETE("/book/:id", handlers.DeleteBook)
	router.PATCH("/book/:id", handlers.UpdateBook)
	router.POST("/collection", handlers.CreateCollection)
	router.PATCH("/collection/:name", handlers.UpdateCollection)
	router.GET("/collections", handlers.GetAllCollection)
	router.DELETE("/collection/:name", handlers.DeleteCollection)
	router.POST("/collection/:name/:id", handlers.CollectBook)
	router.DELETE("/collection/:name/:id", handlers.DiscardBook)

	return router
}
