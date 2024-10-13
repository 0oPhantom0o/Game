package controller

import (
	"game/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *GameController) RequestOtp(c *gin.Context) {
	var user domain.RequestPhone

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := ctrl.Logic.RequestOtp(user.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP requested successfully"})
}
