package structs

type Leaderboard struct {
	Scores []ScoreTotal `json:"scores"`
}

type ScoreTotal struct {
	User string `json:"user"`
	Score int `json:"score"`
}