package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type URLShortener struct {
	urls map[string]string
	mu   sync.Mutex
}

func main() {
	shortener := &URLShortener{
		urls: make(map[string]string),
	}

	http.HandleFunc("/shorten", shortener.HandleShorten)
	http.HandleFunc("/short/", shortener.HandleRedirect)

	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func (u *URLShortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.URL == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := fmt.Sprintf("%d", time.Now().UnixNano())

	u.mu.Lock()
	u.urls[shortURL] = request.URL
	u.mu.Unlock()

	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (u *URLShortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/short/"):]

	u.mu.Lock()
	originalURL, ok := u.urls[shortURL]
	u.mu.Unlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
