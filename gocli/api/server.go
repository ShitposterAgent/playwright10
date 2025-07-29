package api

import (
	"fmt"
	"gocli/types"
	"net/http"
)

func StartServer(cfg *types.Config) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "shitposteragent API is running")
	})
	// Add more endpoints as needed
	http.ListenAndServe(":8080", nil)
}
