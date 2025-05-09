package api_calls

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func playTrack(client *http.Client, trackID, deviceID, token string) error {
	url := buildPlayTrackURL(deviceID)

	body, err := buildPlayTrackBody(trackID)
	if err != nil {
		return err
	}

	req, err := createPlayTrackRequest(url, body, token)
	if err != nil {
		return err
	}

	return sendPlayTrackRequest(client, req)
}

func buildPlayTrackURL(deviceID string) string {
	baseURL := "https://api.spotify.com/v1/me/player/play"
	params := url.Values{}
	params.Add("device_id", deviceID)
	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

func buildPlayTrackBody(trackID string) ([]byte, error) {
	body := map[string]interface{}{
		"uris": []string{"spotify:track:" + trackID},
	}
	return json.Marshal(body)
}

func createPlayTrackRequest(url string, body []byte, token string) (*http.Request, error) {
	req, err := http.NewRequest("PUT", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	set_play_track_headers(req, token)
	return req, nil
}

func sendPlayTrackRequest(client *http.Client, req *http.Request) error {
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func set_play_track_headers(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
}
