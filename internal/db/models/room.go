package models

type Room struct {
	ID    int    `db:"id" json:"id"`
	Code  string `db:"code" json:"code"`
	Score int    `db:"score" json:"score"`
}
