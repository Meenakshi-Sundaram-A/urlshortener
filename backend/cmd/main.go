package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Data struct {
	URL string `json:"urlData"`
}

var urlStore = make(map[string]string)

func generateShortUrl() string {
	const characterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var shortURL = make([]byte, 6)
	for i := range shortURL {
		shortURL[i] = characterSet[r.Intn(len(characterSet))]
	}
	return string(shortURL)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	fmt.Println("The data is:", data.URL)

	shortUrl := generateShortUrl()

	urlStore[shortUrl] = data.URL

	response := map[string]string{
		"shortUrl": fmt.Sprintf("http://localhost:8080/%s", shortUrl),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortKey := params["key"]
	originalURL, exists := urlStore[shortKey]
	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/data", handleData).Methods("POST")
	router.HandleFunc("/{key}", handleRedirect).Methods("GET")
	fmt.Println("Server is running on port 8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, router)

}
