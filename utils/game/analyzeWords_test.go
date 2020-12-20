package game

import (
	"reflect"
	"testing"

	"github.com/daltonscharff/spelling-bee-server/config"
)

func TestCalcPointValue(t *testing.T) {
	word := "immobility"

	expectedPoints := 17
	points := calcPointValue(word)

	if points != expectedPoints {
		t.Errorf("Expected %v, got %v", expectedPoints, points)
	}
}

func TestDefineWord(t *testing.T) {
	conf, err := config.Read("../../config.yaml")
	if err != nil {
		panic(err)
	}

	word := "immobility"

	expected := []definition{
		definition{
			Definition:   "remaining in place",
			PartOfSpeech: "noun",
		},
		definition{
			Definition:   "the quality of not moving",
			PartOfSpeech: "noun",
		},
	}

	definitions := defineWord(word, conf.RapidAPI.Host, conf.RapidAPI.Key)

	if !reflect.DeepEqual(expected, definitions) {
		t.Errorf("Expected %+v, got %+v", expected, definitions)
	}
}

func TestAnalyzeWords(t *testing.T) {
	conf, err := config.Read("../../config.yaml")
	if err != nil {
		panic(err)
	}

	expected := map[string]analyzedWord{
		"immobility": {
			PointValue: 17,
			Definitions: []definition{
				{
					Definition:   "remaining in place",
					PartOfSpeech: "noun",
				},
				{
					Definition:   "the quality of not moving",
					PartOfSpeech: "noun",
				},
			},
		},
	}

	wordMap := analyzeWords([]string{"immobility"}, conf.RapidAPI.Host, conf.RapidAPI.Key)

	if !reflect.DeepEqual(expected, wordMap) {
		t.Errorf("Expected %+v, got %+v", expected, wordMap)
	}
}
