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
		log.Panic("Db didnt init")
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
	//version control
	v1 := Router.Group(constants.Version)

	//authentication routs
	auth := v1.Group(constants.Auth)
	auth.POST(constants.RequestOtp, controller.RequestOtp)
	auth.POST(constants.Submit, controller.GetOtp)
	auth.POST(constants.Login, controller.Login)
	auth.POST(constants.Nickname, controller.NickName)

	//game routs
	game := v1.Group(constants.Game)
	game.GET(constants.RequestQuestion, controller.GameQuestion)
	game.POST(constants.SubmitAnswer, controller.Answer)
	game.GET(constants.ShowTopPlayers, controller.ScoreBoard)

}
