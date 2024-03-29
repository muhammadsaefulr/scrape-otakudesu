package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/otakudesu-scrape/usecase"
)

func GetListAnime(c *gin.Context) {
	animes, err := usecase.GetListAnime()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"animes": animes})

}

func GetAnimeSearch(c *gin.Context) {
	title := c.Query("title")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query Titile Cannot Null"})
		return
	}

	animes, err := usecase.GetAnimeByTitle(title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"animes": animes})

}

func GetAnime(c *gin.Context) {
	title := c.Param("judulAnime")

	animes, err := usecase.GetAnime(title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"animes": animes})
}
