package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/daltonscharff/spelling-bee-server/internal/db"
	"github.com/julienschmidt/httprouter"
)

type Word struct {
	ID           uint   `json:"id"`
	Word         string `json:"word"`
	PuzzleID     uint   `json:"puzzleId"`
	PointValue   uint   `json:"pointValue"`
	Definition   string `json:"definition"`
	PartOfSpeech string `json:"partOfSpeech"`
	Synonym      string `json:"synonym"`
}

func queryWords(query string, args ...interface{}) ([]Word, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	response := []Word{}
	for rows.Next() {
		word := Word{}
		rows.Scan(&word.ID, &word.Word, &word.PuzzleID, &word.PointValue, &word.Definition, &word.PartOfSpeech, &word.Synonym)
		response = append(response, word)
	}

	return response, nil
}

func viewAllWords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, err := queryWords("SELECT * FROM words;")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func viewWord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response, err := queryWords("SELECT * FROM words WHERE id = $1;", p.ByName("id"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch len(response) {
	case 0:
		json.NewEncoder(w).Encode(struct{}{})
		break
	default:
		json.NewEncoder(w).Encode(response[0])
	}
}

func createWord(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var word Word
	json.NewDecoder(r.Body).Decode(&word)

	response, err := queryWords("INSERT INTO words (word, point_value, definition, part_of_speech, synonym, puzzle_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;", word.Word, word.PointValue, word.Definition, word.PartOfSpeech, word.Synonym, word.PuzzleID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch len(response) {
	case 0:
		json.NewEncoder(w).Encode(struct{}{})
		break
	default:
		json.NewEncoder(w).Encode(response[0])
	}
}

func updateWord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var word Word
	json.NewDecoder(r.Body).Decode(&word)

	response, err := queryWords("UPDATE words SET word = $1, point_value = $2, definition = $3, part_of_speech = $4, synonym = $5, puzzle_id = $6 WHERE id = $7 RETURNING *;", word.Word, word.PointValue, word.Definition, word.PartOfSpeech, word.Synonym, word.PuzzleID, p.ByName("id"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch len(response) {
	case 0:
		json.NewEncoder(w).Encode(struct{}{})
		break
	default:
		json.NewEncoder(w).Encode(response[0])
	}
}

func deleteAllWords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var word Word
	json.NewDecoder(r.Body).Decode(&word)

	response, err := queryWords("DELETE FROM words RETURNING *;")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch len(response) {
	case 0:
		json.NewEncoder(w).Encode(struct{}{})
		break
	default:
		json.NewEncoder(w).Encode(response[0])
	}
}

func deleteWord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var word Word
	json.NewDecoder(r.Body).Decode(&word)

	response, err := queryWords("DELETE FROM words WHERE id = $1 RETURNING *;", p.ByName("id"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch len(response) {
	case 0:
		json.NewEncoder(w).Encode(struct{}{})
		break
	default:
		json.NewEncoder(w).Encode(response[0])
	}
}
