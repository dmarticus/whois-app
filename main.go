package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/domainr/whois"
)

func main() {
	// Create a fileServer handler that serves static files
	fileServer := http.FileServer(http.Dir("static/"))

	// tell the http library how we want to handle requests
	// For now, we'll just pass the request along (no error handling yet)
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fileServer.ServeHTTP(w, r)
		},
	)

	// Handle the POST request on /whois
	// The client will send a url-encoded request like:
	//     data=8.8.8.8
	http.HandleFunc("/whois", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Extract the encoded data to perform the whois lookup
		data := r.PostFormValue("data")

		// Perform the query
		result, err := whoisQuery(data)

		// Return a JSON-encoded response
		if err != nil {
			jsonResponse(w, Response{Error: err.Error()})
			return
		}
		jsonResponse(w, Response{Result: result})

	})
	// Finally, start the HTTP server on port 8080.
	// If anything goes wrong, the log.Fatal call will output the error to the console and exit the application.
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type Response struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

func whoisQuery(data string) (string, error) {
	// Run whois on a user-specified object
	response, err := whois.Fetch(data)
	if err != nil {
		return "", err
	}
	return string(response.Body), nil
}

func jsonResponse(w http.ResponseWriter, x interface{}) {
	// JSON-encode some interface x
	bytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	// Write the encoded data to the ResponseWriter.
	// This will send the response to the client.
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
