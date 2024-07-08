package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestOtp(c *gin.Context) {
	var user domain.RequestPhone
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request error")
		return
	}

	err := logic.RequestOtp(user.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "Code is sent")
}
