package api

import (
	"fmt"
	"github.com/aminsaki/golang-clean-web-api/api/routers"
	"github.com/aminsaki/golang-clean-web-api/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("/api")

	v1 := api.Group("v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.TestRoute(test_router)
		routers.Health(health)
	}

	router.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
