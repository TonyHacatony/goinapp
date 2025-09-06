package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	strings = []string{"Hello", "World", "Go", "API"}
	mutex   sync.RWMutex
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mutex.RLock()
	defer mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(strings)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	newString := string(body)
	if newString == "" {
		http.Error(w, "Empty string not allowed", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	strings = append(strings, newString)
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "String '%s' added successfully", newString)
}

func main() {
	http.HandleFunc("/list", getHandler)
	http.HandleFunc("/add", setHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
