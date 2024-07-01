package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOtp(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	err := logic.CompareOtp(user.Phone, user.RandomCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	id, err := logic.GenerateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")

		return
	}
	token, err := logic.GenerateToken(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")

		return
	}
	c.JSON(http.StatusOK, token)

}
