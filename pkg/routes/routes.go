// Package routes contains the layer responsible of mapping all the routes to their respective handlers.
package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joacof98/epicprompts/pkg/prompts"
	"github.com/joacof98/epicprompts/pkg/responses"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()
	router.Get("/", indexHandler)

	// prompts routes
	promptController := prompts.NewPromptController()
	router.Mount("/api/prompts", promptController)

	// responses routes
	responsesController := responses.NewResponseController()
	router.Mount("/api/responses", responsesController)

	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
