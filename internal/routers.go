package internal

import "github.com/gin-gonic/gin"

func (app *MyApp) Routes(r *gin.Engine) {
	r.GET("/", Health)
	{
		r.GET("/person/:person_id/info", app.GetPersonInfo)
		r.POST("/person/create", app.CreatePerson)
	}
}

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"error":   false,
		"message": "check health of the app",
	})
}
