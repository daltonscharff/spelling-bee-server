package router

import (
	"fmt"
	"net/http"

	"github.com/daltonscharff/spelling-bee-server/internal/db"
	"github.com/daltonscharff/spelling-bee-server/internal/models"

	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
)

func getStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := db.Connect()
	if err != nil {
		fmt.Fprintf(w, "Error: could not connect to db")
	}
	defer db.Close()

	puzzle := models.Puzzle{}
	if err := db.QueryRow("SELECT id, date, letters, center, max_score FROM puzzles;").Scan(&puzzle.ID, &puzzle.Date, pq.Array(&puzzle.Letters), &puzzle.Center, &puzzle.MaxScore); err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Puzzle: %+v", puzzle)
}

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/ping", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "Pong\n")
	})
	router.GET("/status", getStatus)
	return router
}
