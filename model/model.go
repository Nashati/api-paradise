package model

type Joke struct {
	Id           int64         `json:"id"`
	Category     string        `json:"category"`
	Body         string        `json:"body"`
	Safe         bool          `json:"safe"`
	LastUpdated  int64         `json:"last_updated"`
	VotesHistory VoteHistories `json:"votes"`
}

type VoteHistories struct {
	Amount  int64         `json:"amount"`
	History []VoteHistory `json:"history"`
}

type VoteHistory struct {
	Username string `json:"username"`
	Up       bool   `json:"up"`
	Date     int64  `json:"date"`
}

type Category struct {
	Id   int64
	Name string
}
