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
		c.JSON(http.StatusBadRequest, "bad request error")
		return
	}
	err := logic.CheckOtp(user.Phone, user.RandomCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, "wronged otp")
		return
	}
	id, err := logic.GenerateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed to create user")
		return
	}
	token, err := logic.GenerateToken(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed to generate token")

		return
	}
	c.JSON(http.StatusOK, token)

}
