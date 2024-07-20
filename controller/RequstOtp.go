package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestOtp(c *gin.Context) {
	var (
		user          domain.RequestPhone
		responseError domain.ErrorHandler
		client        domain.ResponseClient
	)
	err := responseError.Err
	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.Response, err = logic.RequestOtp(user.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": client})
}
