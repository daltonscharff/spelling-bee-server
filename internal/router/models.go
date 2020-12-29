package router

import "time"

type Puzzle struct {
	ID           uint     `json:"-"`
	Date         string   `json:"date"`
	Letters      []string `json:"letters"`
	CenterLetter string   `json:"centerLetter"`
	MaxScore     uint     `json:"maxScore"`
	Words        []Word   `json:"words"`
}

type Word struct {
	ID           uint   `json:"-"`
	Value        string `json:"value"`
	PointValue   uint   `json:"pointValue"`
	Definition   string `json:"definition"`
	PartOfSpeech string `json:"partOfSpeech"`
	Synonym      string `json:"synonym"`
}

type Room struct {
	ID             uint   `json:"-"`
	Code           string `json:"code"`
	Score          string `json:"score"`
	Finds          []Find `json:"finds"`
	WordsRemaining []Word `json:"wordsRemaining"`
}

type Find struct {
	ID         uint      `json:"-"`
	Word       Word      `json:"word"`
	PlayerName string    `json:"playerName"`
	FoundAt    time.Time `json:"foundAt"`
}
