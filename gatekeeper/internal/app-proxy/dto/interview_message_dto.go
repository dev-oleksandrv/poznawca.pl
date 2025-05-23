package dto

type AppInterviewMessageDto struct {
	ID                     string `json:"id"`
	ContentText            string `json:"content_text"`
	ContentTranslationText string `json:"content_translation_text"`
	TipsText               string `json:"tips_text"`
	Role                   string `json:"role"`
	Type                   string `json:"type"`
	CreatedAt              string `json:"created_at"`
}
