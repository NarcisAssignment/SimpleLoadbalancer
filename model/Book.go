package model

type Book struct {
	ID      int64  `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Summary string `json:"summary"`
	Rank    int64  `json:"rank"`
}
