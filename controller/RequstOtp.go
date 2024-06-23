package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestOtp(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	err := logic.TempUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	c.JSON(http.StatusOK, user)
}
