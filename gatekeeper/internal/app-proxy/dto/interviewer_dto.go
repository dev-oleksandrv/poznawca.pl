package dto

type AppInterviewerDto struct {
	ID                        string `json:"id"`
	Name                      string `json:"name"`
	AvatarURL                 string `json:"avatar_url"`
	Description               string `json:"description"`
	DescriptionTranslationKey string `json:"description_translation_key"`
}
