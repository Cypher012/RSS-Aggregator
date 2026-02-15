package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cypher012/rss-aggregator/internal/db"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiConfig.Queries.CreateUser(r.Context(), params.Name)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, user)
}

func (apiConfig *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user db.User) {
	respondWithJSON(w, 200, user)
}
