package utils

import "net/http"

// CORSMiddleware handles headers and automatically cleans up Preflight OPTIONS requests
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Set the headers allowed to cross domains
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Replace with your frontend domain
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// 2. Handle Preflight: If browser sends OPTIONS, return immediately with 200 OK
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 3. Pass down to the actual API endpoint for regular requests (GET, POST, etc.)
		next.ServeHTTP(w, r)
	})
}