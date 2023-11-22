package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vimvek/urlshortner/api"
	"github.com/vimvek/urlshortner/database"
)

func main() {
	store := database.NewStore()
	shortenerAPI := api.NewShortenerAPI(store)

	r := chi.NewRouter()
	shortenerAPI.SetupRoutes(r)

	http.ListenAndServe(":5090", r)
}
