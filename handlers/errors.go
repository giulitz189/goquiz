package handlers

import (
	"net/http"
	"log"
)

func handlerGetQuiz_MarshallingError(rw http.ResponseWriter, err error) {
	log.Printf("GetQuiz Error - Error while marshalling quiz questionaire: %s\n", err)
	rw.WriteHeader(http.StatusInternalServerError)
}

func handlerAnswers_ReadError(rw http.ResponseWriter, err error) {
	log.Printf("Answers Error - Cannot read request body: %s\n", err)
	rw.WriteHeader(http.StatusInternalServerError)
}

func handlerAnswers_UnmarshallingAnsError(rw http.ResponseWriter, err error) {
	log.Printf("Unmarshalling Error - Cannot generate QuizAnswers object: %s\n", err)
	rw.WriteHeader(http.StatusInternalServerError)
}

func handlerAnswers_UnmarshallingSolError(rw http.ResponseWriter, err error) {
	log.Printf("Unmarshalling Error - Cannot create correct answers list: %s\n", err)
	rw.WriteHeader(http.StatusInternalServerError)
}

func handlerAnswers_MarshallingError(rw http.ResponseWriter, err error) {
	log.Printf("Marshalling Error - Cannot marshal Score object: %s\n", err)
	rw.WriteHeader(http.StatusInternalServerError)
}

func handlerLeaderboard_Error(rw http.ResponseWriter, err error) {
	log.Printf("Leaderboard Error - Cannot marshal Leaderboard object: %s\n", err)
	rw.WriteHeader(http.StatusInternalServerError)
}