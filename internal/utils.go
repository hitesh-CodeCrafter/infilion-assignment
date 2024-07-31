package internal

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServerError(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   true,
		"message": err,
	})
}

func CORS() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowMethods("GET,POST")
	return corsConfig
}
