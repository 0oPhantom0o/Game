package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NickName(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	id, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "token has problem")
		return
	}
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	err = logic.NickName(user.NickName, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	c.JSON(http.StatusOK, tokenString)
}
