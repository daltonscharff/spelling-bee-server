package router

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func requireAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if strings.Compare(os.Getenv("SPELLING_BEE_API_KEY"), r.Header.Get("X-SpellingBeeAPI-Key")) != 0 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r, p)
	}
}

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/ping", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "Pong\n")
	})

	router.GET("/api/puzzle", viewPuzzle)
	router.POST("/api/puzzle", requireAuth(createPuzzle))
	// router.PUT("/api/puzzle/:id", requireAuth(updatePuzzle))
	router.DELETE("/api/puzzle/*id", requireAuth(deletePuzzle))

	router.GET("/api/words", viewAllWords)
	router.GET("/api/words/:id", viewWord)
	router.POST("/api/words", requireAuth(createWord))
	router.PUT("/api/words/:id", requireAuth(updateWord))
	router.DELETE("/api/words", requireAuth(deleteAllWords))
	router.DELETE("/api/words/:id", requireAuth(deleteWord))

	// router.GET("/api/rooms/*id", viewRoom)
	// router.POST("/api/rooms", createRoom)
	// router.POST("/api/ws", createWebsocketConnection)

	return router
}
