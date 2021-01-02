package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/daltonscharff/spelling-bee-server/internal/db"

	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
)

func requireAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if strings.Compare(os.Getenv("SPELLING_BEE_API_KEY"), r.Header.Get("X-SpellingBeeAPI-Key")) != 0 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r, ps)
	}
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

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/ping", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "Pong\n")
	})
	router.GET("/api/puzzle", viewPuzzle)
	router.GET("/api/testProtected", requireAuth(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "OK\n")
	}))
	// router.GET("/api/room/:code", viewRoom)
	// router.POST("/api/room", createRoom)
	// router.POST("/api/ws", createWebsocketConnection)
	return router
}
