package dto

import "github.com/go-playground/validator/v10"

type InterviewAIPromptType string

const (
	InterviewAIPromptUserAnswerType InterviewAIPromptType = "user_answer"
	InterviewAIPromptGetResultsType InterviewAIPromptType = "get_results"
)

type InterviewAIPromptBaseDto struct {
	Type InterviewAIPromptType `json:"type" validate:"required"`
}

type InterviewAIPromptUserAnswerDto struct {
	InterviewAIPromptBaseDto
	Content  string `json:"content" validate:"required"`
	Language string `json:"language" validate:"required"`
}

type InterviewAIPromptGetResultsDto struct {
	InterviewAIPromptBaseDto
}

type InterviewAIOutputType string

const (
	InterviewAIOutputQuestionType InterviewAIOutputType = "question"
	InterviewAIOutputResultsType  InterviewAIOutputType = "results"
)

type InterviewAIOutputBaseDto struct {
	Type InterviewAIOutputType `json:"type" validate:"required"`
}

type InterviewAIOutputQuestionDto struct {
	InterviewAIOutputBaseDto
	IsLastMessage          bool   `json:"is_last_message"`
	ContentText            string `json:"content_text" validate:"required"`
	TipsText               string `json:"tips_text"`
	ContentTranslationText string `json:"content_translation_text"`
}

type InterviewAIOutputResultDto struct {
	InterviewAIOutputBaseDto
	GrammarScore     int    `json:"grammar_score"`
	GrammarFeedback  string `json:"grammar_feedback"`
	AccuracyScore    int    `json:"accuracy_score"`
	AccuracyFeedback string `json:"accuracy_feedback"`
	TotalScore       int    `json:"total_score"`
	TotalFeedback    string `json:"translation_text"`
}

type InterviewAICreateThreadInputDto struct {
	EntryMessage         string `json:"entry_message" validate:"required"`
	CharacterDescription string `json:"character_description" validate:"required"`
}

func (d *InterviewAICreateThreadInputDto) Validate() error {
	validate := validator.New()

	return validate.Struct(d)
}

type InterviewAIUserMessageUserAnswerDto struct {
	Content  string `json:"content" validate:"required"`
	Language string `json:"language" validate:"required"`
	ThreadID string `json:"thread_id" validate:"required"`
}

type InterviewAIUserMessageGetResultsDto struct {
	ThreadID string `json:"thread_id" validate:"required"`
}
