package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/otakudesu-scrape/repository"
)

func GetAnime() ([]gin.H, error) {
	return repository.GetAnime()
}

func GetAnimeByTitle(title string) ([]gin.H, error) {
	return repository.GetAnimeByTitle(title)
}
