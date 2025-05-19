package dto

type InterviewResultOutputDto struct {
	GrammarScore     int    `json:"grammar_score"`
	AccuracyScore    int    `json:"accuracy_score"`
	TotalScore       int    `json:"total_score"`
	GrammarFeedback  string `json:"grammar_feedback"`
	AccuracyFeedback string `json:"accuracy_feedback"`
	TotalFeedback    string `json:"total_feedback"`
}
