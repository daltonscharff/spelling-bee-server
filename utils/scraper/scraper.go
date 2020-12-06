package scraper

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/daltonscharff/spelling-bee-server/db"
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

func getLetters(words []string) []byte {
	letters := []byte{}
	allLetters := []byte(strings.Join(words, ""))

	for i := 0; i < len(allLetters); i++ {
		if bytes.Contains(letters, []byte{allLetters[i]}) == false {
			letters = append(letters, allLetters[i])
		}
	}

	return letters
}

func getCenterLetter(words []string, letters []byte) byte {
	isVowel := func(letter byte) bool {
		vowels := []byte("aAeEiIoOuU")
		return bytes.Contains(vowels, []byte{letter})
	}
	letterMap := map[byte]int{}
	var centerLetter byte

	for _, word := range words {
		for _, letter := range letters {
			if bytes.Contains([]byte(word), []byte{letter}) {
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

func Scrape() (db.Puzzle, []db.Word) {
	resp, err := http.Get(sourceURL)
	if err != nil {
		panic(err)
	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		panic(resp.Status)
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	puzzle := db.Puzzle{}
	words := []db.Word{}

	wordList := findWordList(doc)

	for _, word := range wordList {
		words = append(words, db.Word{
			Word: word,
		})
	}

	puzzle.Date = findDate(doc)
	puzzle.Letters = getLetters(wordList)
	puzzle.Center = getCenterLetter(wordList, puzzle.Letters)

	return puzzle, words
}
