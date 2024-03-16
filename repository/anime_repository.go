package repository

import (
	"fmt"

	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func GetAnime() ([]gin.H, error) {
	url := "https://otakudesu.cloud/"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	animes := []gin.H{}
	doc.Find(".venz li").Each(func(i int, s *goquery.Selection) {
		anime := gin.H{
			"episode":    strings.TrimSpace(s.Find(".epz").Text()),
			"dayUpdate":  strings.TrimSpace(s.Find(".epztipe").Text()),
			"dateUpdate": s.Find(".newnime").Text(),
			"url":        s.Find(".thumb a").AttrOr("href", ""),
			"imageUrl":   s.Find(".thumb img").AttrOr("src", ""),
			"title":      strings.TrimSpace(s.Find(".jdlflm").Text()),
		}
		animes = append(animes, anime)
	})

	return animes, nil
}

func GetAnimeByTitle(title string) ([]gin.H, error) {
	url := fmt.Sprintf("https://otakudesu.cloud/?s=%s&post_type=anime", title)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	var animeResults []gin.H

	doc.Find("ul.chivsrc li").Each(func(i int, s *goquery.Selection) {
		// Mendapatkan judul anime
		title := s.Find("h2 a").Text()
		// Mendapatkan link anime
		link, _ := s.Find("h2 a").Attr("href")
		// Mendapatkan thumbnail anime
		thumbnail, _ := s.Find("img").Attr("src")
		// Mendapatkan genre anime
		genres := []string{}
		s.Find(".set b:contains('Genres')").Parent().Find("a").Each(func(i int, genre *goquery.Selection) {
			genres = append(genres, genre.Text())
		})
		// Mendapatkan status anime
		status := s.Find(".set b:contains('Status')").Parent().Text()
		// Mendapatkan rating anime
		rating := s.Find(".set b:contains('Rating')").Parent().Text()

		// Membersihkan rating dari tag dan karakter spasi tambahan
		rating = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(rating, "Rating : "), "\n"))
		status = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(status, "Status : "), "\n"))
		// Menambahkan hasil pencarian anime ke dalam slice
		animeResults = append(animeResults, gin.H{
			"Title":     title,
			"Link":      link,
			"Thumbnail": thumbnail,
			"Genres":    strings.Join(genres, ", "),
			"Status":    status,
			"Rating":    rating,
		})
	})

	return animeResults, nil
}
