package main

import (
	"NOOBDY/questions_server/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "questions-server", log.LstdFlags)

	qh := handlers.NewQuestions(l)

	sm := mux.NewRouter()

	getQuestionRouter := sm.Methods(http.MethodGet).Subrouter()
	getQuestionRouter.HandleFunc("/{id:[0-9]+}", qh.GetQuestion)

	getQuestionsRouter := sm.Methods(http.MethodGet).Subrouter()
	getQuestionsRouter.HandleFunc("/", qh.GetQuestions)

	s := http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal: ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
