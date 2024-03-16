package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/otakudesu-scrape/handlers"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/anime", handlers.GetAnime)
	router.GET("/anime/search", handlers.GetAnimeSearch)
}
