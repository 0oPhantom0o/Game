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
	var userNickname domain.RequestNickName
	if err := c.BindJSON(&userNickname); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	err = logic.UpdateNickName(userNickname.NickName, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	c.JSON(http.StatusOK, "nickName changed to : "+userNickname.NickName)
}
