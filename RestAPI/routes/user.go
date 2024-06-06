package routes

import (
	"fmt"
	"net/http"

	"auth.com/auth/models"
	"auth.com/auth/utils"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	err = user.Save()
	if err != nil {
		fmt.Println(">>>>>>", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to signup"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		fmt.Println(">>>>", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "JWT token generating error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "signup successfully", "user": user, "token": token})

}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		fmt.Println(">>>>", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "JWT token generating error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user logged In", "token": token})

}

func getUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to fetch user record from database"})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func getUser(c *gin.Context) {
	userId := c.GetInt64("userId")
	user, err := models.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to fetch user"})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, user)
}
