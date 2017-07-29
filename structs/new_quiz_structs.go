package structs

type DBQuestions struct {
	Questions []QuestionBlock_DBFormat `json:"questions"`
}

type QuestionBlock_DBFormat struct {
	Question string `json:"question"`
	Score int `json:"score"`
	Answers []string `json:"answers"`
	Correct int `json:"correct"`
}

type Quiz struct {
	ID string `json:"id"`
	Questions []QuestionBlock_SendFormat `json:"questions"`
}

type QuestionBlock_SendFormat struct {
	Question string `json:"question"`
	Answers []string `json:"answers"`
}