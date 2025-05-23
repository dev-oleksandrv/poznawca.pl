package dto

type AppOpenAICreateInterviewRequestDto struct {
	EntryMessage string `json:"entry_message"`
	Description  string `json:"description"`
}

type AppOpenAICreateInterviewResponseDto struct {
	ThreadID string `json:"thread_id"`
}
