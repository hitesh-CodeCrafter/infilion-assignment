package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *MyApp) GetPersonInfo(c *gin.Context) {
	id := c.Param("person_id")

	if id == "" {
		ServerError(c, "id is required")
		return
	}

	data, err := app.GetPersonById(id)
	if err != nil {
		ServerError(c, err.Error())
		return
	}

	c.JSON(200, data)
}

func (app *MyApp) CreatePerson(c *gin.Context) {

	var data Person
	err := c.Bind(&data)
	if err != nil {
		ServerError(c, err.Error())
		return
	}

	err = app.CreatePersonData(data)
	if err != nil {
		ServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Person created successfully",
	})

}
