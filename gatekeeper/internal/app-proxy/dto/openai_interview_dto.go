package dto

type AppOpenAICreateInterviewRequestDto struct {
	EntryMessage string `json:"entry_message"`
	Description  string `json:"description"`
}

type AppOpenAICreateInterviewResponseDto struct {
	ThreadID string `json:"thread_id"`
}

type AppOpenAIInterviewSendUserAnswerRequestDto struct {
	ThreadID string `json:"thread_id" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Language string `json:"language" validate:"required"`
}

type AppOpenAIInterviewGetResultsRequestDto struct {
	ThreadID string `json:"thread_id" validate:"required"`
}

type AppOpenAIInterviewAssistantRequestType string

const (
	AppOpenAIInterviewAssistantRequestUserAnswerType AppOpenAIInterviewAssistantRequestType = "user_answer"
	AppOpenAIInterviewAssistantRequestGetResultsType AppOpenAIInterviewAssistantRequestType = "get_results"
)

type AppOpenAIInterviewAssistantBaseRequestDto struct {
	Type AppOpenAIInterviewAssistantRequestType `json:"type" validate:"required"`
}

type AppOpenAIInterviewAssistantUserAnswerRequestDto struct {
	AppOpenAIInterviewAssistantBaseRequestDto
	Content  string `json:"content" validate:"required"`
	Language string `json:"language" validate:"required"`
}

type AppOpenAIInterviewAssistantGetResultsRequestDto struct {
	AppOpenAIInterviewAssistantBaseRequestDto
}

type AppOpenAIInterviewAssistantResponseType string

const (
	AppOpenAIInterviewAssistantResponseQuestionType AppOpenAIInterviewAssistantResponseType = "question"
	AppOpenAIInterviewAssistantResponseResultsType  AppOpenAIInterviewAssistantResponseType = "results"
)

type AppOpenAIInterviewAssistantBaseResponseDto struct {
	Type AppOpenAIInterviewAssistantResponseType `json:"type" validate:"required"`
}

type AppOpenAIInterviewAssistantQuestionResponseDto struct {
	AppOpenAIInterviewAssistantBaseResponseDto
	IsLastMessage          bool   `json:"is_last_message"`
	ContentText            string `json:"content_text" validate:"required"`
	TipsText               string `json:"tips_text"`
	ContentTranslationText string `json:"content_translation_text"`
}

type AppOpenAIInterviewAssistantResultsResponseDto struct {
	AppOpenAIInterviewAssistantBaseResponseDto
	GrammarScore     int    `json:"grammar_score"`
	GrammarFeedback  string `json:"grammar_feedback"`
	AccuracyScore    int    `json:"accuracy_score"`
	AccuracyFeedback string `json:"accuracy_feedback"`
	TotalScore       int    `json:"total_score"`
	TotalFeedback    string `json:"total_feedback"`
}
