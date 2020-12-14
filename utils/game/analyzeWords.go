package game

type analyzedWord struct {
	Definition   string `json:"definition"`
	PartOfSpeech string `json:"part_of_speech"`
	PointValue   int    `json:"point_value"`
}

func analyzeWords(words []string) (wordMap map[string]analyzedWord, err error) {
	for _, word := range words {
		wordMap[word] = analyzedWord{
			PointValue: calculatePointValue(word),
		}
	}
	return wordMap, nil
}

func calculatePointValue(word string) (points int) {
	points = len(word) - 3

	if len(word) >= 7 {
		letterMap := map[rune]bool{}
		for _, letter := range word {
			letterMap[letter] = true
		}
		uniqueLetters := 0
		for range letterMap {
			uniqueLetters++
		}
		if uniqueLetters == 7 {
			points += 7
		}
	}

	return points
}
