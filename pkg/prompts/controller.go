// This controller is responsible for handling the request and logic for the Prompt entity.
package prompts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joacof98/epicprompts/pkg/initializers"
	"github.com/joacof98/epicprompts/pkg/models"
	"github.com/joacof98/epicprompts/pkg/res"
)

func NewPromptController() http.Handler {
	promptsRouter := chi.NewRouter()
	promptsRouter.Get("/", GetPrompts)
	promptsRouter.Post("/", CreatePrompt)

	return promptsRouter
}

func GetPrompts(w http.ResponseWriter, r *http.Request) {
	var prompts []models.Prompt
	result := initializers.DB.Find(&prompts)
	if result.Error != nil {
		res.ErrorFindingPrompts(w)
	}

	res.RespondWithJSON(w, http.StatusOK, prompts)
}

func CreatePrompt(w http.ResponseWriter, r *http.Request) {
	var prompt PromptDTO
	if err := json.NewDecoder(r.Body).Decode(&prompt); err != nil {
		res.ErrorDecondingJson(w)
		return
	}

	// validations
	if prompt.PromptText == "" || prompt.Category == "" {
		res.ErrorValidPrompt(w)
		return
	} else if prompt.Category != "General" && prompt.Category != "Question" {
		res.ErrorValidCategory(w)
		return
	}

	// check if the prompt already exists
	var promptInBd models.Prompt
	err := initializers.DB.Where("prompt_text = ?", prompt.PromptText).First(&promptInBd).Error
	if err != nil {
		// prompt is new, insert it to DB
		newPrompt := models.Prompt{Category: prompt.Category, PromptText: prompt.PromptText}
		result := initializers.DB.Create(&newPrompt)
		if result.Error != nil {
			res.ErrorCreatingPrompt(w, result.Error.Error())
			return
		}

		res.RespondWithJSON(w, http.StatusOK, "Prompt created successfully")
	} else {
		res.ErrorExistingPrompt(w)
	}
}
