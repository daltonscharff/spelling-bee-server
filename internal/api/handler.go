package api

import (
	"net/http"
	"os"
	"strings"

	"github.com/daltonscharff/spelling-bee-server/internal/postgres"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	*httprouter.Router
	store postgres.Store
}

func requireAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if strings.Compare(os.Getenv("SPELLING_BEE_API_KEY"), r.Header.Get("X-SpellingBeeAPI-Key")) != 0 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r, p)
	}
}

func NewHandler(store postgres.Store) *Handler {
	h := &Handler{
		Router: httprouter.New(),
		store:  store,
	}
	h.GET("/api/words/:id", viewWord(h))
	h.GET("/api/words", viewAllWords(h))
	h.POST("/api/words", requireAuth(createWord(h)))
	h.PUT("/api/words/:id", requireAuth(updateWord(h)))
	h.DELETE("/api/words/:id", requireAuth(deleteWord(h)))
	h.DELETE("/api/words", requireAuth(deleteAllWords(h)))

	// router.GET("/api/puzzle", viewPuzzle)
	// router.POST("/api/puzzle", requireAuth(createPuzzle))
	// router.PUT("/api/puzzle/:id", requireAuth(updatePuzzle))
	// router.DELETE("/api/puzzle/*id", requireAuth(deletePuzzle))

	// router.GET("/api/words", viewAllWords)
	// router.GET("/api/words/:id", viewWord)
	// router.POST("/api/words", requireAuth(createWord))
	// router.PUT("/api/words/:id", requireAuth(updateWord))
	// router.DELETE("/api/words", requireAuth(deleteAllWords))
	// router.DELETE("/api/words/:id", requireAuth(deleteWord))

	return h
}
