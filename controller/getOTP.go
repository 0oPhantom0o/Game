package controller

import (
	"game/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *GameController) GetOtp(c *gin.Context) {

	var user domain.RequestRandomCode

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := ctrl.Logic.GenerateUser(user.Phone, user.RandomCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
