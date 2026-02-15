package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Cypher012/rss-aggregator/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB      *pgxpool.Pool
	Queries *db.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the enviornment")
	}

	pool := NewPool()
	defer pool.Close()

	queries := db.New(pool)

	log.Println("Database connected!")

	apiCfg := apiConfig{
		DB:      pool,
		Queries: queries,
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	r.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)
}
