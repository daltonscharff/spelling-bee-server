package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/daltonscharff/spelling-bee-server/internal/db"
	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
)

type Puzzle struct {
	ID           uint     `json:"-"`
	Date         string   `json:"date"`
	Letters      []string `json:"letters"`
	CenterLetter string   `json:"centerLetter"`
	MaxScore     uint     `json:"maxScore"`
	Words        []Word   `json:"words"`
}

func viewPuzzle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := db.Connect()
	if err != nil {
		fmt.Fprintf(w, "Error: could not connect to db")
	}
	defer db.Close()

	response := Puzzle{}
	if err := db.QueryRow("SELECT id, date, letters, center, max_score FROM puzzles ORDER BY date DESC LIMIT 1;").Scan(&response.ID, &response.Date, pq.Array(&response.Letters), &response.CenterLetter, &response.MaxScore); err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT word, point_value FROM words WHERE puzzle_id = $1;", &response.ID)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		word := Word{}
		rows.Scan(&word.Word, &word.PointValue)
		response.Words = append(response.Words, word)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createPuzzle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func updatePuzzle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func deletePuzzle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
