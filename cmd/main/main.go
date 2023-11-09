package main

import (
	"net/http"

	"github.com/mlgpuntnl/workout-generator/internal/routes"

	"github.com/joho/godotenv"
)

const port = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	router := routes.BuildRouter()

	// Start server
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
