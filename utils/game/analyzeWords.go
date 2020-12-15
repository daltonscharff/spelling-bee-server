package game

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type analyzedWord struct {
	PointValue  int          `json="pointValue"`
	Definitions []definition `json="definitions"`
}

type definition struct {
	Definition   string `json="definition"`
	PartOfSpeech string `json="partOfSpeech"`
}

func analyzeWords(words []string, APIHost string, APIKey string) (wordMap map[string]analyzedWord) {
	for _, word := range words {
		wordMap[word] = analyzedWord{
			PointValue:  calculatePointValue(word),
			Definitions: defineWord(word, APIHost, APIKey),
		}
	}
	return wordMap
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

func defineWord(word string, APIHost string, APIKey string) []definition {
	req, _ := http.NewRequest("GET", "https://wordsapiv1.p.rapidapi.com/words/"+word+"/definitions", nil)

	req.Header.Add("x-rapidapi-key", APIKey)
	req.Header.Add("x-rapidapi-host", APIHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	resObj := struct {
		Definitions []definition `json:"definitions"`
	}{}

	if err = json.Unmarshal(body, &resObj); err != nil {
		panic(err)
	}

	return resObj.Definitions
}
