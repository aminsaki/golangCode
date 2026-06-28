package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type header struct {
	UserId  string
	Browser string
}
type personData struct {
	FirstName    string `json:"first_name" binding:"required,min=4"`
	LastName     string `json:"last_name"  binding:"required,min=4"`
	MobileNumber string `json:"mobile_number"  binding:"required,mobile,len=11"` // اصلاح شد
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Hello World",
	})
}
func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "users",
	})
}

func (h *TestHandler) UsersById(c *gin.Context) {

	id := c.Param("id")
	c.JSON(200, gin.H{
		"result": "usersById",
		"id":     id,
	})
}

func (h *TestHandler) UserByUsername(c *gin.Context) {
	username := c.Param("username")
	c.JSON(200, gin.H{
		"result":   "username",
		"username": username,
	})
}

func (h *TestHandler) Accounts(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"result": "Accounts",
		"id":     id,
	})
}
func (h *TestHandler) AddUser(c *gin.Context) {

	id := c.Param("id")
	c.JSON(200, gin.H{
		"result": "AddUser",
		"id":     id,
	})
}

func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("userId")
	c.JSON(200, gin.H{
		"result": "headerBinder1",
		"userId": userId,
	})
}
func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	err := c.BindHeader(&header)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"result": "headerBinder2",
		"userId": header,
	})
}

func (h *TestHandler) FormBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBind(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "FormBinder",
		"person": p,
	})
}
