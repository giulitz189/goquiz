package handlers

import (
	"net/http"
	"io/ioutil"
	"sync"
	"encoding/json"

	"goquiz/structs"
	"github.com/go-chi/chi"
)

var (
	sended           []string
	leaderboard      structs.Leaderboard

	sendedMutex      sync.RWMutex
	leaderboardMutex sync.RWMutex
)

func loadQuestions(dbquiz *structs.DBQuestions) error {
	qdata, err1 := ioutil.ReadFile("/Utenti/Giuliano/gowork/src/goquiz/handlers/questions.json")
	if err1 != nil {
		return err1
	}

	if err2 := json.Unmarshal(qdata, &dbquiz); err2 != nil {
		return err2
	}
	return nil
}

func HandlerGetQuiz(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var quiz structs.Quiz
	var ok bool
	quiz.ID, ok = getRequestID(r.Context())
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var dbq structs.DBQuestions
	if err1 := loadQuestions(&dbq); err1 != nil {
		handlerGetQuiz_MarshallingError(rw, err1)
		return
	}
	for _, q := range dbq.Questions {
		var qblock structs.QuestionBlock_SendFormat
		qblock.Question = q.Question
		qblock.Answers = q.Answers
		quiz.Questions = append(quiz.Questions, qblock)
	}

	raw, err2 := json.Marshal(quiz)
	if err2 != nil {
		handlerGetQuiz_MarshallingError(rw, err2)
		return
	}

	sendedMutex.Lock()
	sended = append(sended, quiz.ID)
	sendedMutex.Unlock()

	rw.WriteHeader(http.StatusOK)
	rw.Write(raw)
}

func HandlerAnswers(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := chi.URLParam(r, "quizId")
	if id == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var idFound = false
	sendedMutex.RLock()
	for _, acceptedID := range sended {
		if acceptedID == id {
			idFound = true
			break
		}
	}
	sendedMutex.RUnlock()
	if !idFound {
		rw.WriteHeader(http.StatusForbidden)
		return
	}

	raw_r, err_r := ioutil.ReadAll(r.Body)
	if err_r != nil {
		handlerAnswers_ReadError(rw, err_r)
		return
	}

	var quizAns structs.QuizAnswers
	if err1 := json.Unmarshal(raw_r, &quizAns); err1 != nil {
		handlerAnswers_UnmarshallingAnsError(rw, err1)
		return
	}

	var dbq structs.DBQuestions
	if err2 := loadQuestions(&dbq); err2 != nil {
		handlerAnswers_UnmarshallingSolError(rw, err2)
		return
	}

	var score structs.QuizScore
	score.Score = 0
	for idx, answer := range quizAns.Answers {
		var result structs.Result
		result.Given = answer
		result.Correct = dbq.Questions[idx].Correct
		if result.Given == result.Correct {
			score.Score += dbq.Questions[idx].Score
		}
		score.Results = append(score.Results, result)
	}

	raw_w, err_w := json.Marshal(score)
	if err_w != nil {
		handlerAnswers_MarshallingError(rw, err_w)
		return
	}

	var lbscore structs.ScoreTotal
	lbscore.User = id
	lbscore.Score = score.Score
	leaderboardMutex.Lock()
	leaderboard.Scores = append(leaderboard.Scores, lbscore)
	leaderboardMutex.Unlock()

	rw.WriteHeader(http.StatusOK)
	rw.Write(raw_w)
}

func HandlerLeaderboard(rw http.ResponseWriter, r *http.Request) {
	leaderboardMutex.RLock()
	defer leaderboardMutex.RUnlock()

	raw, err := json.Marshal(leaderboard)
	if err != nil {
		handlerLeaderboard_Error(rw, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(raw)
	r.Body.Close()
}