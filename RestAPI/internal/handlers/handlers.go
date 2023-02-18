package Handler

import (
	"fmt"
	Logic "myapp/internal/logic"
	"net/http"

	Model "myapp/internal/Model"

	"github.com/gin-gonic/gin"
)

func PostPerson(c *gin.Context) {
	var newPerson Model.Person
	newPerson.Email = c.Request.FormValue("email")
	newPerson.Phone = c.Request.FormValue("phone")
	newPerson.FirstName = c.Request.FormValue("firstName")
	newPerson.LastName = c.Request.FormValue("lastName")
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Logic.Create(newPerson)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func GetPersons(c *gin.Context) {
	persons, err := Logic.Read()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	c.IndentedJSON(http.StatusOK, persons)
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	persons, err := Logic.ReadOne(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	c.IndentedJSON(http.StatusOK, persons)
}


func DeleteById(c *gin.Context) {
	id := c.Param("id")
	err := Logic.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	c.IndentedJSON(http.StatusOK, "Запись удалена")
}

