package api_calls

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strings"
)

func Get_Redirect(clientID string) (fullURL, verifier string, err error) {
	verifier, err = generateVerifier()
	if err != nil {
		return "", "", err
	}

	challenge := generateChallenge(verifier)

	params := buildRedirectParams(clientID, challenge)

	baseURL := "https://accounts.spotify.com/authorize"
	fullURL = fmt.Sprintf("%s?%s", baseURL, params.Encode())

	return fullURL, verifier, nil
}

func generateVerifier() (string, error) {
	codeVerifier := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, codeVerifier)
	if err != nil {
		return "", err
	}
	verifier := base64.URLEncoding.EncodeToString(codeVerifier)
	return strings.TrimRight(verifier, "="), nil
}

func generateChallenge(verifier string) string {
	hash := sha256.New()
	hash.Write([]byte(verifier))
	challenge := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return strings.TrimRight(challenge, "=")
}

func buildRedirectParams(clientID, challenge string) url.Values {
	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("response_type", "code")
	params.Add("redirect_uri", "http://localhost:8080/hit-of-the-day")
	params.Add("code_challenge_method", "S256")
	params.Add("code_challenge", challenge)
	params.Add("scope", "user-read-playback-state user-modify-playback-state user-read-private")
	return params
}
