package model

type Genre struct {
	Id    int16  `json:"id" db:"genre_id"`
	Genre string `json:"genre" db:"genre"`
}
