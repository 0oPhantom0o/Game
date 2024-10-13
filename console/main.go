package main

import (
	"context"
	"game/app"
	"game/constants"
	"game/controller"
	"game/logic"
	"game/repository"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	if err := app.InitDb(); err != nil {
		log.Panicf("DataBase is not running:%v", err)
	}

}
func main() {
	mongodb, err := app.Collection()
	if err != nil {
		log.Fatalf("Error initializing databases: %v", err)
	}
	ctx := context.Background()
	redisdb := app.RedisDB
	repo := repository.NewMongoRepository(redisdb, mongodb, ctx)
	svc := logic.NewRestaurantService(repo)
	ctrl := controller.NewGameController(svc)

	//router Init
	router := gin.Default()
	setupRoute(router, ctrl)
	//Run Server
	err = router.Run(constants.ServerPort)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func setupRoute(Router *gin.Engine, ctrl *controller.GameController) {
	//setup routes
	//version control
	v1 := Router.Group(constants.Version)

	//authentication routs
	auth := v1.Group(constants.Auth)
	auth.POST(constants.RequestOtp, ctrl.RequestOtp)
	auth.POST(constants.Submit, ctrl.GetOtp)
	auth.POST(constants.Login, ctrl.Login)
	auth.POST(constants.Nickname, ctrl.NickName)
	//game routs
	game := v1.Group(constants.Game)
	game.GET(constants.RequestQuestion, ctrl.RequestQuestion)
	game.POST(constants.SubmitAnswer, ctrl.Answer)
	game.GET(constants.ShowTopPlayers, ctrl.ScoreBoard)

}
