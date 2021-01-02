package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Questions handler for getting questions
type Questions struct {
	l *log.Logger
}

// NewQuestions returns a new questions handler with the given logger
func NewQuestions(l *log.Logger) *Questions {
	return &Questions{l}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

func getQuestionID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
