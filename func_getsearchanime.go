package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func searhanime() {
	r := gin.Default()

	r.GET("/anime", func(c *gin.Context) {
		search := c.Query("search")

		if search == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
			return
		}

		resp, err := http.Get("https://otakudesu.cloud/?s=" + search + "&post_type=anime")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		pattern := `<h2><a href="(.*?)"[^>]*>(.*?)<\/a><\/h2>.*?<b>Genres<\/b> : (.*?)<\/div>.*?<b>Status<\/b> : (.*?)<\/div>.*?<b>Rating<\/b> : (.*?)<\/div>`
		re := regexp.MustCompile(pattern)

		matches := re.FindAllSubmatch(body, -1)

		// Buat slice untuk menyimpan hasil
		animes := make([]gin.H, 0)

		// Loop melalui setiap hasil pencocokan dan tambahkan ke slice
		for _, match := range matches {
			link := string(match[1])
			title := string(match[2])
			genre := string(match[3])
			status := string(match[4])
			rating := string(match[5])

			reTag := regexp.MustCompile(`<[^>]*>`)
			cleanGenre := reTag.ReplaceAllString(genre, "")

			genreArray := strings.Split(cleanGenre, ", ")

			if strings.Contains(strings.ToLower(title), strings.ToLower(search)) {
				animes = append(animes, gin.H{"link": link, "title": title, "genre": genreArray, "status:": status, "rating": rating})
			}
		}

		c.JSON(http.StatusOK, gin.H{"animes": animes})
	})

	r.Run(":8080")
}
