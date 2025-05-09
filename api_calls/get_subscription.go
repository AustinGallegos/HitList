package api_calls

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get_Subscription(client *http.Client, token string) (string, error) {
	req, err := buildUserRequest(token)
	if err != nil {
		return "", err
	}

	resp, err := sendUserRequest(client, req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := readResponseBody(resp)
	if err != nil {
		return "", err
	}

	return decodeSubscription(body)
}

func buildUserRequest(token string) (*http.Request, error) {
	url := "https://api.spotify.com/v1/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	set_subscription_headers(req, token)
	return req, nil
}

func sendUserRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return resp, nil
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status: %d, body: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func decodeSubscription(body []byte) (string, error) {
	var userResp struct {
		Subscription string `json:"product"`
	}
	err := json.Unmarshal(body, &userResp)
	if err != nil {
		return "", fmt.Errorf("error decoding JSON: %w", err)
	}
	return userResp.Subscription, nil
}

func set_subscription_headers(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}
