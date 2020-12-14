package newGame

import (
	"sort"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var validHTML string = `<div id="date-and-pic"> <h2>Tuesday, December 1, 2020</h2> <div id="bee-pic"> <img src="pics/20201201.png" width="200"> <!--width="350"--> </div> <div><p><a href="Bee_20201130.html">Answers to yesterday's bee</a></p></div> </div> <div id="main-answer-list" class="answer-list" style="display: block;"> <ul class="column-list"> <li> blot </li> <li> blotto </li> <li> bolt </li> <li> boot </li> <li> booty </li> <li> bottom </li> <li> <mark><strong>immobility</strong></mark> </li> <li> itty </li> <li> lilt </li> <li> limit </li> <li> lobotomy </li> <li> loot </li> <li> lotto </li> <li> mitt </li> <li> <mark><strong>mobility</strong></mark> </li> <li> molt </li> <li> moot </li> <li> motility </li> <li> motto </li> <li> obit </li> <li> omit </li> <li> till </li> <li> tilt </li> <li> toil </li> <li> toll </li> <li> tomb </li> <li> tomboy </li> <li> tomtit </li> <li> tool </li> <li> toot </li> </ul> </div>`

var expectedDate string = "2020-12-01"
var expectedLetters []string = []string{"b", "i", "l", "m", "o", "t", "y"}
var expectedCenter string = "t"
var expectedWords []string = []string{
	"blot",
	"blotto",
	"bolt",
	"boot",
	"booty",
	"bottom",
	"immobility",
	"itty",
	"lilt",
	"limit",
	"lobotomy",
	"loot",
	"lotto",
	"mitt",
	"mobility",
	"molt",
	"moot",
	"motility",
	"motto",
	"obit",
	"omit",
	"till",
	"tilt",
	"toil",
	"toll",
	"tomb",
	"tomboy",
	"tomtit",
	"tool",
	"toot",
}

func TestFindDate(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(validHTML))
	if err != nil {
		panic(err)
	}

	date := findDate(doc)

	if strings.Compare(date, expectedDate) != 0 {
		t.Errorf("Expected %s, got %s", expectedDate, date)
	}
}

func TestFindWordList(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(validHTML))
	if err != nil {
		panic(err)
	}

	wordList := findWordList(doc)
	sort.Strings(wordList)

	for i, word := range wordList {
		if strings.Compare(word, expectedWords[i]) != 0 {
			t.Errorf("Expected %s, got %s", expectedWords, wordList)
		}
	}
}

func TestGetLetters(t *testing.T) {
	letters := getLetters(expectedWords)

	for i, letter := range letters {
		if letter != letters[i] {
			t.Errorf("Expected %v, got %v", expectedLetters, letters)
		}
	}
}

func TestGetCenterLetter(t *testing.T) {
	centerLetter := getCenterLetter(expectedWords, expectedLetters)

	if centerLetter != expectedCenter {
		t.Errorf("Expected %v, got %v", expectedCenter, centerLetter)
	}
}
