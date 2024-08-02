package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handleReset(w http.ResponseWriter, r *http.Request) {
	cfg.mu.Lock()
	cfg.fileserverHits = 0
	fmt.Println("Hits reset to 0")
	cfg.mu.Unlock()
	w.WriteHeader(200)
}
