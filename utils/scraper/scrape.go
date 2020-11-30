package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const sourceURL = "https://nytbee.com"

func findDate(doc *goquery.Document) *string {
	const dateLayout = "Monday, January 2, 2006"

	text := doc.Find("#date-and-pic h2").First().Text()

	dt, err := time.Parse(dateLayout, text)
	if err != nil {
		panic(err)
	}

	date := dt.Format("2006-01-02")
	return &date
}

func findWordList(doc *goquery.Document) *[]string {
	words := []string{}

	doc.Find("#main-answer-list .column-list li").Each(func(i int, s *goquery.Selection) {
		word := strings.TrimSpace(s.Text())
		word = strings.ToLower(word)
		words = append(words, word)
	})

	return &words
}

func createLetterMap(words []string) *map[rune]int {
	letterMap := map[rune]int{}
	allLetters := strings.Join(words, "")

	for _, char := range allLetters {
		letterMap[char]++
	}

	return &letterMap
}

func getLetters(letterMap *map[rune]int) *[]rune {
	letters := []rune{}

	for key := range *letterMap {
		letters = append(letters, key)
	}

	return &letters
}

func getCenterLetter(letterMap *map[rune]int) rune {
	isVowel := func(letter rune) bool {
		vowels := "aAeEiIoOuU"
		return strings.ContainsRune(vowels, letter)
	}
	max := struct {
		character rune
		amount    int
	}{0, -1}

	for k, v := range *letterMap {
		if v > max.amount || (v == max.amount && !isVowel(k)) {
			max.character = k
			max.amount = v
		}
	}

	return max.character
}

type GameData struct {
	Date         string   `json:"gameDate"`
	Words        []string `json:"words"`
	Letters      []string `json:"letters"`
	CenterLetter string   `json:"centerLetter"`
}

func (g *GameData) JSON() string {
	data, err := json.Marshal(g)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func Scrape() *GameData {
	resp, err := http.Get(sourceURL)
	if err != nil {
		panic(err)
	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		panic(resp.Status)
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	data := GameData{}
	data.Date = *findDate(doc)
	data.Words = *findWordList(doc)

	letterMap := *createLetterMap(data.Words)
	data.Letters = strings.Split(string(*getLetters(&letterMap)), "")
	data.CenterLetter = string(getCenterLetter(&letterMap))

	return &data
}
