package main

import (
	"katze/logger"
	"katze/src/graphql"
	"net/http"
)

func main() {

	handler, err := graphql.GraphqlHandler()
	if err != nil {
		logger.Errorf("Error creating graphql handler: %v", err)
	}
	http.Handle("/graphql", handler)
	http.ListenAndServe(":80", nil)
}
