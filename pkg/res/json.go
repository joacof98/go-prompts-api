// Package res includes functions for responding to HTTP requests with serialized JSONs.
package res

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

// Custom responses for Prompts
func ErrorDecondingJson(w http.ResponseWriter) {
	RespondWithError(w, http.StatusInternalServerError, "Error parsing request body. Please check the type of the fields provided")
}

func ErrorValidPrompt(w http.ResponseWriter) {
	RespondWithError(w, http.StatusBadRequest, "Please provide a valid prompt (text and category must be strings)")
}

func ErrorFindingPrompts(w http.ResponseWriter) {
	RespondWithError(w, http.StatusBadRequest, "Cannot find prompts")
}

func ErrorValidCategory(w http.ResponseWriter) {
	RespondWithError(w, http.StatusBadRequest, "Please provide a valid category (only 'General' and 'Question' are allowed)")
}

func ErrorCreatingPrompt(w http.ResponseWriter, err string) {
	RespondWithError(w, http.StatusBadRequest, "Error creating new prompt: "+err)
}

func ErrorExistingPrompt(w http.ResponseWriter) {
	RespondWithError(w, http.StatusBadRequest, "That prompt already exists. Please use another one")
}

// Custom responses for Responses
func ErrorFindingResponses(w http.ResponseWriter) {
	RespondWithError(w, http.StatusBadRequest, "Cannot find responses")
}

func ErrorValidResponse(w http.ResponseWriter) {
	RespondWithError(w, http.StatusBadRequest, "Please provide a valid response (text and category must be strings)")
}

func ErrorValidPromptId(w http.ResponseWriter) {
	RespondWithError(w, http.StatusBadRequest, "Please provide a valid prompt id")
}

func ErrorCreatingResponse(w http.ResponseWriter, err string) {
	RespondWithError(w, http.StatusBadRequest, "Error creating new response for the prompt: "+err)
}
