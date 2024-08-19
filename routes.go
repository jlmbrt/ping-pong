package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - Request from %s\n", time.Now().Format(time.RFC822), r.RemoteAddr)

	contentType, response := buildResponse(r.Header.Get("Accept"), "pong")

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(200)
	w.Write([]byte(response))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - Request from %s\n", time.Now().Format(time.RFC822), r.RemoteAddr)

	msg := fmt.Sprintf("Hello there 🙋\nI'm %s, %s", config.Name, config.Message)
	contentType, response := buildResponse(r.Header.Get("Accept"), msg)

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(200)
	w.Write([]byte(response))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func sleepHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - Request from %s\n", time.Now().Format(time.RFC822), r.RemoteAddr)

	inputTime, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/sleep/"), 10, 8)

	if err != nil {
		inputTime = 1
	}

	duration := time.Duration(inputTime) * time.Second

	fmt.Printf("Sleeping during %ds\n", inputTime)

	time.Sleep(duration)

	w.WriteHeader(200)
}

func buildResponse(acceptHeader string, msg string) (string, string) {

	accept := strings.Split(acceptHeader, ",")[0]

	var response string
	var contentType string

	switch accept {
	case "application/json":
		contentType = accept
		response = fmt.Sprintf("{\"server\":\"%s\",\"response\":\"%s\"}", config.Name, msg)
	case "text/html":
		contentType = accept
		response = fmt.Sprintf(`
<html>
<head><meta charset="UTF-8"></head>
<body>
<h1>I'm %s</h1><h2>And I say: %s</h2>
</body>
</html>
		`, config.Name, msg)
	default:
		contentType = "plain/text"
		response = msg
	}

	return contentType, response
}
