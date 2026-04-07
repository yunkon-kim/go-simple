package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/yunkon-kim/go-simple/internal/ip"
)

func main() {
	// Page handler - serves the main HTML page
	http.HandleFunc("/", handleIndex)

	// API endpoint - returns the visitor's IP as JSON
	http.HandleFunc("/api/my-ip", handleMyIP)

	// API endpoint - returns the server's public IP as JSON
	http.HandleFunc("/api/server-ip", handleServerIP)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleIndex renders the main HTML page.
func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// handleMyIP returns the visitor's IP address.
func handleMyIP(w http.ResponseWriter, r *http.Request) {
	clientIP := ip.GetClientIP(r)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"ip": clientIP,
	})
}

// handleServerIP returns the server's public IP address.
func handleServerIP(w http.ResponseWriter, r *http.Request) {
	publicIP, err := ip.GetPublicIP()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to get server IP",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"ip": publicIP,
	})
}
