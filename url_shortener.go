package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

func GetAllURLs() map[string]string {
	content, err := os.ReadFile("shortened-urls.json")

	if err != nil {
		return make(map[string]string)
	}

	// Parse the JSON content
	urlsMap := make(map[string]string)

	err = json.Unmarshal(content, &urlsMap)

	if err != nil {
		return make(map[string]string)
	}

	return urlsMap
}

func SaveURLToFile(hashedUrl, originalUrl string) {
	alreadySavedURLs := GetAllURLs()
	alreadySavedURLs[hashedUrl] = originalUrl

	// Convert the map to JSON
	jsonContent, err := json.Marshal(alreadySavedURLs)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("shortened-urls.json", jsonContent, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func HashURL(urlToShorten string) string {
	randomStr := randomString(8)
	combined := urlToShorten + randomStr
	hash := sha256.Sum256([]byte(combined))
	encoded := base64.URLEncoding.EncodeToString(hash[:])

	return encoded[:8] // Take the first 8 characters for a shorter URL
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
