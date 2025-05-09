package api_calls

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Get_Token(client *http.Client, clientID, code, verifier string) (string, error) {
	req, err := buildTokenRequest(clientID, code, verifier)
	if err != nil {
		return "", err
	}

	resp, err := sendTokenRequest(client, req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := readTokenResponse(resp)
	if err != nil {
		return "", err
	}

	token, err := parseTokenJSON(body)
	if err != nil {
		return "", err
	}

	return token, nil
}

func buildTokenRequest(clientID, code, verifier string) (*http.Request, error) {
	baseURL := "https://accounts.spotify.com/api/token"

	params := url.Values{}
	params.Set("grant_type", "authorization_code")
	params.Set("code", code)
	params.Set("redirect_uri", "http://localhost:8080/hit-of-the-day")
	params.Set("client_id", clientID)
	params.Set("code_verifier", verifier)

	req, err := http.NewRequest("POST", baseURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to build token request: %w", err)
	}

	set_token_headers(req)
	return req, nil
}

func sendTokenRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send token request: %w", err)
	}
	return resp, nil
}

func readTokenResponse(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading token response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status: %d, body: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func parseTokenJSON(body []byte) (string, error) {
	var tokenResp struct {
		AccessToken string `json:"access_token"`
	}

	err := json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", fmt.Errorf("error decoding token JSON: %w", err)
	}

	return tokenResp.AccessToken, nil
}

func set_token_headers(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}
