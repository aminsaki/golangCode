package api

import (
	"fmt"
	"github.com/aminsaki/golang-clean-web-api/api/middlewares"
	"github.com/aminsaki/golang-clean-web-api/api/routers"
	"github.com/aminsaki/golang-clean-web-api/api/validations"
	"github.com/aminsaki/golang-clean-web-api/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config) {
	router := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IrananMobileNumberValiudator)
	}

	router.Use(middlewares.Cors(cfg))
	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("/api")

	v1 := api.Group("v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.TestRoute(test_router)
		routers.Health(health)
	}
	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}

	router.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
