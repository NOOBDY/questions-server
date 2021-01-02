package handlers

import (
	"NOOBDY/questions_server/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetQuestions handlers GET request and returns all current products
func (q *Questions) GetQuestions(w http.ResponseWriter, r *http.Request) {
	q.l.Println("Handle GET Questions")

	questions := data.GetQuestions()

	err := data.ToJSON(questions, w)

	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// GetQuestion handles GET requests
func (q *Questions) GetQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
	}

	q.l.Println("Handle GET Question", id)

	question, err := data.GetQuestion(id)

	switch err {
	case nil:
	case data.ErrQuestionNotFound:
		q.l.Println("[ERROR] fetching question", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	default:
		q.l.Println("[ERROR] fetching question", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	err = data.ToJSON(question, w)
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
	}
}
