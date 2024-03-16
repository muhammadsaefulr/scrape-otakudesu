package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/otakudesu-scrape/repository"
)

func GetListAnime() ([]gin.H, error) {
	return repository.GetListAnime()
}

func GetAnimeByTitle(title string) ([]gin.H, error) {
	return repository.GetAnimeByTitle(title)
}
