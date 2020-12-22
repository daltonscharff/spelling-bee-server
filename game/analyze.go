package game

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type analyzedWord struct {
	PointValue  int          `json="pointValue"`
	Definitions []definition `json="definitions"`
}

type definition struct {
	Definition   string `json="definition"`
	PartOfSpeech string `json="partOfSpeech"`
}

func calcPointValue(word string) (points int) {
	if len(word) < 4 {
		points = 0
	} else if len(word) == 4 {
		points = 1
	} else {
		points = len(word)
	}

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

func defineWord(word string) []definition {
	req, _ := http.NewRequest("GET", "https://wordsapiv1.p.rapidapi.com/words/"+word+"/definitions", nil)

	req.Header.Add("x-rapidapi-key", os.Getenv("RAPID_API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPID_API_HOST"))

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

func analyzeWords(words []string) map[string]analyzedWord {
	wordMap := map[string]analyzedWord{}
	for _, word := range words {
		wordMap[word] = analyzedWord{
			PointValue:  calcPointValue(word),
			Definitions: defineWord(word),
		}
	}
	return wordMap
}
