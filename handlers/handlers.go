package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiURL = "https://api.usw2.pure.cloud/api/v2/architect/prompts?pageSize=220&pageNumber=1"
)

type Response struct {
	Entities []Entity `json:"entities"`
}

type Entity struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Resources   []Resource `json:"resources"`
}

type Resource struct {
	Language  string `json:"language"`
	MediaURI  string `json:"mediaUri"`
	TTSString string `json:"ttsString"`
}

func GetPrompt(accessToken string) ([]Entity, error) {
	// Create HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return response.Entities, nil
}

func ExtractFields(responseData []Resource) ([]Resource, error) {
	if len(responseData) == 0 {
		return []Resource{}, fmt.Errorf("no resources available")
	}
	return responseData, nil
}
