package middlewares

import (
	"github.com/aminsaki/golang-clean-web-api/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", cfg.Cors.AllowOrigin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Max-Age", "21600")
		c.Set("content-type", "application/json")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()

	}

}
