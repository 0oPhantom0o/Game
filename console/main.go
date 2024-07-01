package main

import (
	"game/app"
	"game/constants"
	"game/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := app.InitDb(); err != nil {
		log.Panic()
	}

	//Router Init
	Router := gin.Default()
	setupRoute(Router)

	//Run Server
	err := Router.Run(constants.ServerPort)
	if err != nil {
		return
	}
}

func setupRoute(Router *gin.Engine) {
	//setup routes
	Router.POST("/sign_up", controller.RequestOtp)
	Router.POST("/sign_in", controller.GetOtp)
	Router.POST("/Auth", controller.NickName)
	Router.GET("/Calculate", controller.Calculate)
	Router.POST("/result", controller.Result)
	Router.GET("/ScoreBoard/:number", controller.ScoreBoard)

}
