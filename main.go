package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/extism/go-pdk"
)

// API endpoint
const apiEndpoint = "https://cleanuri.com/api/v1/shorten"

type requestPayload struct {
	URL string `json:"url"`
}

type responsePayload struct {
	ResultURL string `json:"result_url"`
	Error     string `json:"error"`
}

//export run
func run() int32 {
	// Read input URL from plugin input
	longURL := strings.TrimSpace(string(pdk.Input()))

	// Validate the long URL
	if strings.ContainsAny(longURL, " \t\n") {
		pdk.SetError(fmt.Errorf("Invalid URL: URL should not contain spaces, tabs, or newlines"))
		return 1
	}

	// URL-encode the long URL
	encodedURL := &requestPayload{URL: longURL}
	payloadBytes, err := json.Marshal(encodedURL)
	if err != nil {
		pdk.SetError(fmt.Errorf("Failed to create request payload: %v", err))
		return 1
	}

	// Create HTTP request
	req := pdk.NewHTTPRequest(pdk.MethodPost, apiEndpoint)
	req.SetHeader("Content-Type", "application/json")
	req.SetBody(bytes.NewBuffer(payloadBytes).Bytes())

	// Send the HTTP request
	res := req.Send()

	// Parse response
	if res.Status() != 200 {
		pdk.SetError(fmt.Errorf("API request failed with status: %v", res.Status()))
		return 1
	}

	var response responsePayload
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		pdk.SetError(fmt.Errorf("Failed to parse API response: %v", err))
		return 1
	}

	// Check for errors in the response
	if response.Error != "" {
		pdk.SetError(fmt.Errorf("API error: %v", response.Error))
		return 1
	}

	// Set the output to the shortened URL
	pdk.OutputString(response.ResultURL)
	return 0
}

func main() {}
