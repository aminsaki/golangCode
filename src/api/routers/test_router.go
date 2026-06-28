package routers

import (
	"github.com/aminsaki/golang-clean-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRoute(router *gin.RouterGroup) {
	h := handlers.NewTestHandler()
	router.GET("/", h.Test)
	router.GET("/users", h.Users)
	router.GET("/users/:id", h.UsersById)
	router.GET("/users/get-user-by-username/:username", h.UserByUsername)
	router.GET("/users/:id/accounts", h.Accounts)
	router.POST("/add-user", h.AddUser)
	///header
	router.POST("/binder/header1", h.HeaderBinder1)
	router.POST("/binder/header2", h.HeaderBinder2)
	//query
	router.POST("/binder/body", h.FormBinder)

}
