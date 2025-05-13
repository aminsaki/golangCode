package routers

import (
	"github.com/aminsaki/golang-clean-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRoute(router *gin.RouterGroup) {
	h := handlers.NewTestHandler()
	router.GET("/", h.Test)
}
