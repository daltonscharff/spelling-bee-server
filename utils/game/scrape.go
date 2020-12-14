package game

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const sourceURL = "https://nytbee.com"

func findDate(doc *goquery.Document) string {
	const inputLayout = "Monday, January 2, 2006"
	const outputLayout = "2006-01-02"

	text := doc.Find("#date-and-pic h2").First().Text()

	dt, err := time.Parse(inputLayout, text)
	if err != nil {
		panic(err)
	}

	date := dt.Format(outputLayout)
	return date
}

func findWordList(doc *goquery.Document) []string {
	words := []string{}

	doc.Find("#main-answer-list .column-list li").Each(func(i int, s *goquery.Selection) {
		word := strings.TrimSpace(s.Text())
		word = strings.ToLower(word)
		words = append(words, word)
	})

	return words
}

func getLetters(words []string) []string {
	letters := []byte{}
	allLetters := strings.Join(words, "")

	for _, letter := range allLetters {
		if strings.ContainsRune(string(letters), letter) == false {
			letters = append(letters, byte(letter))
		}
	}

	return strings.Split(string(letters[:]), "")
}

func getCenterLetter(words []string, letters []string) string {
	isVowel := func(letter string) bool {
		vowels := "aAeEiIoOuU"
		return strings.Contains(vowels, letter)
	}
	letterMap := map[string]int{}
	var centerLetter string

	for _, word := range words {
		for _, letter := range letters {
			if strings.Contains(word, letter) {
				letterMap[letter]++
			}
		}
	}

	for k, v := range letterMap {
		if v == len(words) && (centerLetter == "" || isVowel(centerLetter)) {
			centerLetter = k
		}
	}

	return centerLetter
}

func scrape() (date string, letters []string, centerLetter string, words []string) {
	resp, err := http.Get(sourceURL)
	if err != nil {
		panic(err)
	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		panic(resp.Status)
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	words = findWordList(doc)
	date = findDate(doc)
	letters = getLetters(words)
	centerLetter = getCenterLetter(words, letters)

	return date, letters, centerLetter, words
}
