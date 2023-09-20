// This controller is responsible for handling the request and logic for the Response entity.
package responses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joacof98/epicprompts/pkg/initializers"
	"github.com/joacof98/epicprompts/pkg/models"
	"github.com/joacof98/epicprompts/pkg/res"
)

func NewResponseController() http.Handler {
	responsesRouter := chi.NewRouter()
	responsesRouter.Get("/", getResponses)
	responsesRouter.Post("/", createResponse)
	responsesRouter.Get("/{promptId}", getResponsesByPromptId)

	return responsesRouter
}

func getResponses(w http.ResponseWriter, r *http.Request) {
	var responses []models.Response
	result := initializers.DB.Find(&responses)
	if result.Error != nil {
		res.ErrorFindingResponses(w)
	}

	res.RespondWithJSON(w, http.StatusOK, responses)
}

func getResponsesByPromptId(w http.ResponseWriter, r *http.Request) {
	promptId := chi.URLParam(r, "promptId")

	id, err := strconv.Atoi(promptId)
	if err != nil {
		res.ErrorValidPromptId(w)
		return
	}

	var responses []models.Response
	result := initializers.DB.Where("prompt_id = ?", id).Find(&responses)
	if result.Error != nil {
		res.ErrorFindingResponses(w)
	}

	res.RespondWithJSON(w, http.StatusOK, responses)
}

func createResponse(w http.ResponseWriter, r *http.Request) {
	var responseDto ResponseDTO
	if err := json.NewDecoder(r.Body).Decode(&responseDto); err != nil {
		res.ErrorDecondingJson(w)
		return
	}

	// validations
	if responseDto.ResponseText == "" || responseDto.Category == "" {
		res.ErrorValidResponse(w)
		return
	} else if responseDto.PromptID == 0 {
		res.ErrorValidPromptId(w)
		return
	} else if responseDto.Category != "General" && responseDto.Category != "Question" {
		res.ErrorValidCategory(w)
		return
	}

	newResponse := models.Response{ResponseText: responseDto.ResponseText, Category: responseDto.Category, PromptID: responseDto.PromptID}
	result := initializers.DB.Create(&newResponse)
	if result.Error != nil {
		res.ErrorCreatingResponse(w, result.Error.Error())
		return
	}

	res.RespondWithJSON(w, http.StatusCreated, "Response for the prompt created")
}
