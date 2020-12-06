package scraper

import (
	"sort"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

const validHTML string = `<div id="date-and-pic"> <h2>Tuesday, December 1, 2020</h2> <div id="bee-pic"> <img src="pics/20201201.png" width="200"> <!--width="350"--> </div> <div><p><a href="Bee_20201130.html">Answers to yesterday's bee</a></p></div> </div> <div id="main-answer-list" class="answer-list" style="display: block;"> <ul class="column-list"> <li> blot </li> <li> blotto </li> <li> bolt </li> <li> boot </li> <li> booty </li> <li> bottom </li> <li> <mark><strong>immobility</strong></mark> </li> <li> itty </li> <li> lilt </li> <li> limit </li> <li> lobotomy </li> <li> loot </li> <li> lotto </li> <li> mitt </li> <li> <mark><strong>mobility</strong></mark> </li> <li> molt </li> <li> moot </li> <li> motility </li> <li> motto </li> <li> obit </li> <li> omit </li> <li> till </li> <li> tilt </li> <li> toil </li> <li> toll </li> <li> tomb </li> <li> tomboy </li> <li> tomtit </li> <li> tool </li> <li> toot </li> </ul> </div>`

var expectedGameData GameData = GameData{
	Date: "2020-12-01",
	Words: []string{
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
	},
	Letters: []byte{
		'b',
		'i',
		'l',
		'm',
		'o',
		't',
		'y',
	},
	CenterLetter: 't',
}

func TestFindDate(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(validHTML))
	if err != nil {
		panic(err)
	}

	date := findDate(doc)

	if strings.Compare(date, expectedGameData.Date) != 0 {
		t.Errorf("Expected %s, got %s", expectedGameData.Date, date)
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
		if strings.Compare(word, expectedGameData.Words[i]) != 0 {
			t.Errorf("Expected %s, got %s", expectedGameData.Words, wordList)
		}
	}
}

func TestGetLetters(t *testing.T) {
	letters := getLetters(expectedGameData.Words)

	for i, letter := range letters {
		if letter != letters[i] {
			t.Errorf("Expected %U, got %U", expectedGameData.Letters, letters)
		}
	}
}

func TestGetCenterLetter(t *testing.T) {
	centerLetter := getCenterLetter(expectedGameData.Words, expectedGameData.Letters)

	expectedCenterLetter := expectedGameData.CenterLetter

	if centerLetter != expectedCenterLetter {
		t.Errorf("Expected %U, got %U", expectedCenterLetter, centerLetter)
	}
}
