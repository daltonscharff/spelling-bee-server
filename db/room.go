package db

type Room struct {
	ID   int    `json:"id,omitempty"`
	Code string `json:"code"`
}
