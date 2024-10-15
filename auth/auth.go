package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	tokenURL = "https://login.usw2.pure.cloud/oauth/token"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Error       string `json:"error,omitempty"`
}

func GetAccessToken(clientID, clientSecret string) (*TokenResponse, error) {
	// Encode client ID and secret
	auth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	// Create request body
	body := []byte("grant_type=client_credentials")

	// Create HTTP request
	req, err := http.NewRequest("POST", tokenURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+auth)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var tokenResponse TokenResponse
	err = json.Unmarshal(respBody, &tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if tokenResponse.Error != "" {
		return nil, fmt.Errorf("error from server: %s", tokenResponse.Error)
	}

	return &tokenResponse, nil
}
