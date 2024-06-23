package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, "error")
		return
	}
	Id, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "error")
		return
	}
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	user.ID, err = logic.ConvertStringToPrimivite(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	err = logic.NickName(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	c.JSON(http.StatusOK, tokenString)
}
