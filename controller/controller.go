package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)


type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type ApiDataAlbums struct {
	Items []struct {
		Name         string `json:"name"`
		Release_date string `json:"release_date"`
		Total_tracks int    `json:"total_tracks"`
		Images       []struct {
			Url string `json:"url"`
		} `json:"images"`
	} `json:"items"`
}


func getToken() (string, error) {
	clientID := "61357de565a24c3c8e0cc0eb25411589"
	clientSecret := "c55f966b457a419e9671189afcd4d2eb"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var token TokenResponse
	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}


func Damso(w http.ResponseWriter, r *http.Request) {
	token, err := getToken()
	if err != nil {
		http.Error(w, "Error getting token", http.StatusInternalServerError)
		return
	}

	apiURL := "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums"
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error calling Spotify API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	var decodeData ApiDataAlbums
	err = json.Unmarshal(body, &decodeData)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(decodeData)
}
