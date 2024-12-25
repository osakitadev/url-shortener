package main

import (
	"log"
	"net/http"
	"strings"
)

func redirectURL(w http.ResponseWriter, r *http.Request) {
	urlPath := r.PathValue("hashurl")

	// For some reason, the browser is requesting the favicon.ico when I visit
	// the root path 🤷‍♂️
	if strings.HasPrefix(urlPath, "favicon.ico") {
		http.NotFound(w, r)
		return
	}

	allURLs := GetAllURLs()
	targetURL, ok := allURLs[urlPath]

	if !ok {
		http.NotFound(w, r)
		return
	}

	log.Println("Redirecting to", targetURL)
	http.Redirect(w, r, targetURL, http.StatusSeeOther)
}

func ServeServer() {
	http.HandleFunc("/{hashurl}", redirectURL)
	http.ListenAndServe(":8080", nil)
}
