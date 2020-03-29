package router

import (

	"awesomeProject/controllers"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func MapRoutes() {

	/*****************************ROUTERS*********************************/

	/*******************************GET***********************************/
	go router.GET("/getcharacters", controllers.GetCharacters)
	go router.GET("/getcharacters/:id", controllers.GetCharactersID)
	go router.GET("/getepisodes", controllers.GetEpisodes)
	go router.GET("/getepisodes/:id", controllers.GetEpisodesID)
	go router.GET("/getquotes", controllers.GetQuotes)
	go router.GET("/getquotes/:id", controllers.GetQuoutesID)
	go router.GET("/getdeaths", controllers.GetDeaths)
	/*****************************POST************************************/
	go router.POST("/createcharacters", controllers.PostCharacters)
	/*****************************DELETE**********************************/
	go router.DELETE("/deletecharacters/:id", controllers.DeleteCharactersID)
	/******************************PUT*************************************/
	go router.PUT("/updatecharacters", controllers.PutCharacters)
}

func Run() {
	router.Run(":8086")
}
