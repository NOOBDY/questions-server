package handlers

import (
	"NOOBDY/questions_server/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Questions struct {
	l *log.Logger
}

func NewQuestions(l *log.Logger) *Questions {
	return &Questions{l}
}

func (q *Questions) GetQuestions(w http.ResponseWriter, r *http.Request) {
	q.l.Println("Handle GET Questions")

	lq := data.GetQuestions()

	err := lq.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (q *Questions) GetQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
	}

	q.l.Println("Handle GET Question", id)

	question := data.GetQuestion(id)

	err = question.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
	}
}
