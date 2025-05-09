package api_calls

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	// "github.com/joho/godotenv"
)

type SearchItems struct {
	ID string `json:"id"`
}

type SearchResponse struct {
	Tracks struct {
		Items []SearchItems `json:"items"`
	} `json:"tracks"`
}

func SearchTrack(query, token string) ([]SearchItems, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track&limit=10", url.QueryEscape(query))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build SearchSong request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling SearchSong API request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var searchResp SearchResponse
	json.Unmarshal(body, &searchResp)
	return searchResp.Tracks.Items, nil
}
