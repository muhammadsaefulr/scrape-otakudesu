package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/otakudesu-scrape/routes"
)

func main() {
	r := gin.Default()

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	panic(err)
	// }

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Service Is UP"})
	})

	routes.InitRoutes(r)

	r.Run(":8000")
}
