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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token has problem"})
		return
	}
	var user domain.RequestNickName
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err = logic.UpdateNickName(user.NickName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "you changed nickname before"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"nickName changed to": user.NickName})
}
