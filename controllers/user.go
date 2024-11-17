package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, user)
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
