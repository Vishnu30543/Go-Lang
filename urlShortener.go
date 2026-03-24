package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"time"
	"net/http"
	"encoding/json"
)

type URL struct {
	Id		  string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	CreationDate  time.Time `json:"creation_date"`
}

/*
	d9785234 -> {
					Id: d9785234,
					OriginalURL: https://www.google.com,
					ShortURL: http://localhost:3000/redirect/d9785234,
					CreationDate: 2024-06-01
				}
*/

var urlDB = make(map[string]URL)

// Generates a short URL by hashing the original URL using MD5 and taking the first 8 characters of the hash.
func generateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL))	// It converts the original URL into a byte slice and writes it to the hasher
	data := hasher.Sum(nil)				// It computes the hash and returns the resulting byte slice
	fmt.Println("Hasher Sum: ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("Hash: ", hash)
	fmt.Println("Hash: ", hash[:8])		// We are taking only the first 8 characters of the hash to create a shorter URL
	return hash[:8]
}

// Creates a short URL for the given original URL and stores it in the urlDB map.
func createShortURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL
	urlDB[id] = URL{
		Id: id,
		OriginalURL: originalURL,
		ShortURL: "http://localhost:3000/redirect/" + shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

// Retrieves the original URL based on the short URL ID.
// It checks if the ID exists in the urlDB map and returns the corresponding URL struct and a boolean indicating whether it was found or not.
func getURL(id string) (URL, error) {
	url, exists := urlDB[id]
	if !exists {
		return URL{}, fmt.Errorf("URL not found")
	}
	return url, nil
}

func RootPageURLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Root page is called...")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		OriginalURL string `json:"original_url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Original URL received:", data.OriginalURL)

	shortURL_ := createShortURL(data.OriginalURL)
	// Direct Print
	// fmt.Fprintln(w, "Short URL: ", shortURL)

	// JSON Response
	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: "http://localhost:3000/redirect/" + shortURL_,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):] // Extract the short URL ID from the path
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}


func main() {
	fmt.Println("Learning URL Shortener in Golang...")
	// originalURL := "https://github.com/Vishnu30543"
	// shortURL := generateShortURL(originalURL)
	// fmt.Println("Short URL: ", shortURL)


	http.HandleFunc("/", RootPageURLHandler)
	http.HandleFunc("/short", ShortURLHandler)
	http.HandleFunc("/redirect/", RedirectURL)

	fmt.Println("Starting port on 3000...")
	
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

}