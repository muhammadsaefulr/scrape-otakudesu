package models

type Anime struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Title  string `json:"title"`
	Rating string `json:"rating"`
}
