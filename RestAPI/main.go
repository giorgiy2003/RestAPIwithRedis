package main

import (
	Handler "myapp/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/person", Handler.GetPersons)
	router.GET("/person/:id", Handler.GetById)
	router.POST("/person", Handler.PostPerson)
	router.DELETE("/person/:id", Handler.DeleteById)
	_ = router.Run("localhost:8080")
}
