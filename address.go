package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GetExternalIP fetches the external IP address from the provided IP service endpoint
func GetExternalIP(ipServiceEndpoint string) (string, error) {
	// Use a public IP service to get the external IP address
	resp, err := http.Get(ipServiceEndpoint)
	if err != nil {
		return "", fmt.Errorf("failed to get external IP: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Strip any new lines or extra whitespace from the IP address
	return strings.TrimSpace(string(ip)), nil
}
