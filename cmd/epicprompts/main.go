package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joacof98/epicprompts/pkg/initializers"
	"github.com/joacof98/epicprompts/pkg/routes"
)

func init() {
	// initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	router := routes.NewRouter()

	port := os.Getenv("PORT")
	addr := "localhost:" + port
	log.Printf("Server is running on %s\n", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
