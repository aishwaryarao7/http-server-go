package main

import (
	"net"
	"strings"
	"testing"
)

// Unit tests
func TestGetURLHeaders(t *testing.T) {
	// Test case 1: Valid URL
	url := "https://example.com"
	headers := getURLHeaders(url)
	if len(headers) == 0 {
		t.Errorf("Expected non-empty headers, got empty headers")
	}

	// Test case 2: Invalid URL
	url = "https://invalidurl"
	headers = getURLHeaders(url)
	if len(headers) != 0 {
		t.Errorf("Expected empty headers, got non-empty headers")
	}
}

func TestGetUserAgent(t *testing.T) {
	// Mock net.Conn
	conn := &net.TCPConn{}

	// Test case 1: Valid response
	expectedUserAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	getUserAgent(conn)
	if !strings.Contains(output, expectedUserAgent) {
		t.Errorf("Expected user agent %q, got %q", expectedUserAgent, output)
	}

	// Test case 2: Invalid response
	expectedUserAgent = ""
	getUserAgent(conn)
	if !strings.Contains(output, expectedUserAgent) {
		t.Errorf("Expected user agent %q, got %q", expectedUserAgent, output)
	}
}

func TestMain(t *testing.T) {
	// Mock net.Listener
	l := &net.TCPListener{}

	// Test case 1: Successful connection
	conn := &net.TCPConn{}
	main()
	if !strings.Contains(output, "Accepted new Connection") {
		t.Errorf("Expected output to contain 'Accepted new Connection', got %q", output)
	}

	// Test case 2: Error accepting connection
	conn = nil
	main()
	if !strings.Contains(output, "Error accepting connection") {
		t.Errorf("Expected output to contain 'Error accepting connection', got %q", output)
	}
}
