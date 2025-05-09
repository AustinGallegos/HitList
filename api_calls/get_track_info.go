package api_calls

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TrackResponse struct {
	Name  string `json:"name"`
	Album struct {
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	} `json:"album"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
}

func GetTrackInfo(client *http.Client, track_id, token string) (map[string]string, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/tracks/%s", track_id)

	req, err := prepareRequest(url, token)
	if err != nil {
		return nil, err
	}

	resp, err := sendRequest(client, req)
	if err != nil {
		return nil, err
	}

	trackResp, err := parseTrackResponse(resp.Body)
	if err != nil {
		return nil, err
	}

	return mapTrackData(trackResp), nil
}

func prepareRequest(url, token string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	setCoverHeaders(req, token)
	return req, nil
}

func sendRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	return resp, nil
}

func parseTrackResponse(body io.Reader) (*TrackResponse, error) {
	var trackResp TrackResponse
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&trackResp)
	if err != nil {
		return nil, err
	}
	return &trackResp, nil
}

func mapTrackData(trackResp *TrackResponse) map[string]string {
	return map[string]string{
		"artistName": trackResp.Artists[0].Name,
		"trackName":  trackResp.Name,
		"imageLink":  trackResp.Album.Images[0].URL,
	}
}

func setCoverHeaders(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}
