package structs

type QuizAnswers struct {
	Answers []int `json:"answers"`
}

type QuizScore struct {
	Score int `json:"score"`
	Results []Result `json:"results"`
}

type Result struct {
	Given int `json:"given"`
	Correct int `json:"questions"`
}