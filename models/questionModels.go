package models

type Answer struct {
	CorrectAnswer   string   `json:"correct_answer"`
	IncorrectAnswer []string `json:"incorrect_answer"`
}

type Question struct {
	Category        string   `json:"category"`
	ID              string   `json:"id"`
	CorrectAnswer   string   `json:"correctAnswer"`
	IncorrectAnswer []string `json:"incorrectAnswers"`
	Question        string   `json:"question"`
	Tags            []string `json:"tags"`
	Type            string   `json:"type"`
	Difficulty      string   `json:"difficulty"`
	Regions         []string `json:"regions"`
	IsNiche         bool     `json:"isNiche"`
}
