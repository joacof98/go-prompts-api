package responses

type ResponseDTO struct {
	ResponseText string `json:"text"`
	Category     string `json:"category"`
	PromptID     uint   `json:"prompt_id"`
}
