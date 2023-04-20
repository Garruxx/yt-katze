package main

import (
	"fmt"
	"katze/logger"
	"katze/src/graphql"
	"net/http"
	"os"
)

func main() {
	// Set up port
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	addr := fmt.Sprintf(":%s", port)

	// Graphql handler
	handler, err := graphql.GraphqlHandler()
	if err != nil {
		logger.Errorf("Error creating graphql handler: %v", err)
	}
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			handler.ServeHTTP(w, r)
		}()
	})

	// Start server
	fmt.Println("Server started on port 80")
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Errorf("Error starting server: %v", err)
	}
}
