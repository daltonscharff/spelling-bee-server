package router

type Room struct {
	ID             uint   `json:"-"`
	Code           string `json:"code"`
	Score          string `json:"score"`
	Finds          []Find `json:"finds"`
	WordsRemaining []Word `json:"wordsRemaining"`
}
