package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	internalMiddleware "github.com/GabrielFMPinheiro/rate-limiter-full-cycle/api/internal_middleware"
	"github.com/GabrielFMPinheiro/rate-limiter-full-cycle/infra/cache"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	cache := cache.NewCache()

	limitRequestPerSecondDefault, err := strconv.ParseInt(os.Getenv("LIMIT_REQUEST_PER_SECOND_DEFAULT"), 10, 64)

	if err != nil {
		limitRequestPerSecondDefault = 10
	}

	rateLimiter := internalMiddleware.NewRateLimiterMiddleware(cache, limitRequestPerSecondDefault, "api-key.json")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(rateLimiter.Middleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
