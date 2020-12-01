package scraper

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const sourceURL = "https://nytbee.com"

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

func getLetters(words []string) []rune {
	letters := []rune{}
	letterMap := map[rune]int{}
	allLetters := strings.Join(words, "")

	for _, char := range allLetters {
		letterMap[char]++
	}

	for key := range letterMap {
		letters = append(letters, key)
	}

	return letters
}

func getCenterLetter(words []string, letters []rune) rune {
	isVowel := func(letter rune) bool {
		vowels := "aAeEiIoOuU"
		return strings.ContainsRune(vowels, letter)
	}
	letterMap := map[rune]int{}
	var centerLetter rune

	for _, word := range words {
		for _, letter := range letters {
			if strings.ContainsRune(word, letter) {
				letterMap[letter]++
			}
		}
	}

	for k, v := range letterMap {
		if v == len(words) && (centerLetter == 0 || isVowel(centerLetter)) {
			centerLetter = k
		}
	}

	return centerLetter
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

	data := GameData{
		Date:  findDate(doc),
		Words: findWordList(doc),
	}

	letters := getLetters(data.Words)
	data.Letters = strings.Split(string(letters), "")
	data.CenterLetter = string(getCenterLetter(data.Words, letters))

	return &data
}
