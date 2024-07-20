package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

	var client domain.RequestPhone
	var resToken domain.ResponseToken
	if err := c.BindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	var err error
	resToken.Token, err = logic.Login(client.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, resToken)
}
