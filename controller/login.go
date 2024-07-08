package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

	var client domain.RequestPhone
	if err := c.BindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request error")
		return
	}
	token, err := logic.Login(client.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, token)
}
