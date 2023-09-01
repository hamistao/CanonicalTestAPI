package routes

import (
	"canonicalTestAPI/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.POST("/book", handlers.CreateBook)
	router.GET("/books", handlers.Query)
	router.DELETE("/book/{id}", handlers.DeleteBook)
	router.PATCH("/book/{id}", handlers.UpdateBook)
	router.POST("/collection", handlers.CreateCollection)
	router.PATCH("/collection/{name}", handlers.UpdateCollection)
	router.GET("/collections", handlers.UpdateCollection)
	router.DELETE("/collection/{name}", handlers.DeleteCollection)
	router.POST("/collection/{name}/{id}", handlers.CollectBook)
	router.PATCH("/collection/{name}/{id}", handlers.CollectBook)

	return router
}
