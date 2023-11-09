package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func BuildRouter() *mux.Router {
	r := mux.NewRouter()

	// Handle routes
	r.HandleFunc("/test", test)

	return r
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
