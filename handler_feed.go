package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cypher012/rss-aggregator/internal/db"
)

func (apiConfig *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiConfig.Queries.CreateFeed(r.Context(), db.CreateFeedParams{
		Url:    params.Url,
		Name:   params.Name,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, feed)
}

func (apiConfig *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feed, err := apiConfig.Queries.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed: %v", err))
		return
	}

	respondWithJSON(w, 200, feed)
}
