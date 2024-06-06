package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"auth.com/auth/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Param id"})
		fmt.Println(err)
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "unable to fetch event"})
		fmt.Println(err)
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not register event"})
		fmt.Println(err)
		return
	}

	fmt.Println(">>>>>")
	context.JSON(http.StatusCreated, gin.H{"message": "Event Registered"})

}

func cancelRegistration(context *gin.Context) {

}
