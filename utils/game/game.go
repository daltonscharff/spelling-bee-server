package game

func Refresh() {
	date, letters, centerLetter, words := scrape()
	wordMap, _ := analyzeWords(words)
	puzzleID, err := updateDB(date, letters, centerLetter, wordMap)
}
