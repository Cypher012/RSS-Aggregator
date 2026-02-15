package main

import (
	"fmt"
	"net/http"

	"github.com/Cypher012/rss-aggregator/internal/auth"
	"github.com/Cypher012/rss-aggregator/internal/db"
)

type authHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error %v", err))
		}

		user, err := apiCfg.Queries.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
