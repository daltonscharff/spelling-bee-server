package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func viewAllWords(h *Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		response, err := h.store.Words()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func viewWord(h *Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id, err := strconv.ParseUint(p.ByName("id"), 0, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response, err := h.store.Word(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func createWord(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func updateWord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// func deleteAllWords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

// }

func deleteWord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
