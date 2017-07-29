package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"goquiz/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(logRequest)
	// Il frontend utilizzato per questo progetto utilizza comandi OPTIONS al posto di GET
	r.Options("/new-quiz", handlers.HandlerGetQuiz)
	r.Options("/leaderboard", handlers.HandlerLeaderboard)
	r.Route("/submit", func (r chi.Router) {
		r.Post("/{quizId}", handlers.HandlerAnswers)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

const logTimeFormat = "02/01/2006 15:04:05"

func logRequest (h http.Handler) http.Handler {
	logger := log.New(os.Stdout, "", 0)

	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r.WithContext(handlers.SetRequestID(r.Context())))
		end := time.Now()

		logger.Printf(
			"[%s] %s -- %s %s -- %s\n",
			start.Format(logTimeFormat),
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			end.Sub(start),
		)
	})
}