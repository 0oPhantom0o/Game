package controller

import (
	"game/constants"
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestOtp(c *gin.Context) {
	var user domain.RequestPhone
	var responseError domain.ErrorHandler
	var client domain.ResponseClient

	err := responseError.Err
	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = logic.RequestOtp(user.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client.Response = constants.CodeIsSent + user.Phone
	c.JSON(http.StatusOK, gin.H{"message": client})
}

//a
