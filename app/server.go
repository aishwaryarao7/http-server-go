package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func getURLHeaders(url string) map[string]interface{} {
	response, err := http.Head(url)
	if err != nil {
		log.Fatal("Error: Unable to download URL (", url, ") with error: ", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatal("Error: HTTP Status = ", response.Status)
	}

	headers := make(map[string]interface{})

	for k, v := range response.Header {
		headers[strings.ToLower(k)] = string(v[0])
	}

	return headers
}

func getUserAgent(conn net.Conn) {
	headers := getURLHeaders("/user-agent")
	fmt.Println(headers)
	_, err := conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\n" + headers["user-agent"].(string) + "\r\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Accepted new Connection")

	buf := make([]byte, 4096)
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(string(buf), "/user-agent") {
		// Respond with a 200 OK response and content type of text/plain
		getUserAgent(conn)
		return
	}
	lines := strings.Split(string(buf), "\r\n")
	path := strings.Fields(lines[0])[1]

	switch {
	case path == "/":
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	case strings.HasPrefix(path, "/user-agent"):
		{
			getUserAgent(conn)
		}
	case strings.HasPrefix(path, "/echo"):
		{
			text := strings.TrimPrefix(path, "/echo/")
			size := fmt.Sprint(len(text))
			_, err = conn.Write([]byte("HTTP/1.1 200 OK \r\nContent-Type: text/plain\r\nContent-Length: " + size + "\r\n\r\n" + text + "\r\n"))
		}
	default:
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

}
