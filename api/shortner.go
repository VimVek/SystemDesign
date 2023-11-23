package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/vimvek/urlshortner/database"
	"github.com/vimvek/urlshortner/hash"
	"github.com/vimvek/urlshortner/models"
)

// ShortenerAPI represents the URL shortener API.
type ShortenerAPI struct {
	store *database.Store
	mu    sync.RWMutex
}

// NewShortenerAPI creates a new ShortenerAPI.
func NewShortenerAPI(store *database.Store) *ShortenerAPI {
	return &ShortenerAPI{
		store: store,
	}
}

func (api *ShortenerAPI) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var input struct {
		LongURL string `json:"longUrl"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//hashe function to generate the short URL.
	shortURL := hash.Hashe(input.LongURL)
	api.store.SaveURL(models.URL{ShortURL: shortURL, LongURL: input.LongURL})

	response := map[string]string{"shortURL": shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (api *ShortenerAPI) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	longURL, err := api.store.GetURL(shortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	response := map[string]string{"longURL": longURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	// http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}
func (api *ShortenerAPI) GetAllURLs(w http.ResponseWriter, r *http.Request) {
	api.mu.RLock()
	defer api.mu.RUnlock()

	allURLs := make(map[string]string)
	for shortURL, longURL := range api.store.GetAllURLs() {
		allURLs[shortURL] = longURL
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allURLs)
}

// API routes.
func (api *ShortenerAPI) SetupRoutes(r chi.Router) {
	r.Post("/api/v1/data/shorten", api.CreateShortURL)
	r.Get("/api/v1/{shortURL}", api.RedirectURL)
	r.Get("/api/v1/all", api.GetAllURLs)
}
