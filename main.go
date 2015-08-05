package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc(`/{filename:.*\.md}`, ServeGithubMarkdown).Methods("GET")
	r.HandleFunc(`/{filename:.*}`, ServeFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":9999", r))
}

func ServeGithubMarkdown(response http.ResponseWriter, request *http.Request) {
	filename := mux.Vars(request)["filename"]

	markdown, err := ioutil.ReadFile(filename)
	if err != nil {
		http.Error(response, "Failed to read markdown file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	body := map[string]string{
		"text": string(markdown),
		"mode": "gfm",
	}
	b, err := json.Marshal(body)
	if err != nil {
		http.Error(response, "Failed to build Github API request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	r, _ := http.NewRequest("POST", "https://api.github.com/markdown", bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		http.Error(response, "Error connecting to Github API: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != 200 {
		http.Error(response, "Error connecting to Github API: "+resp.Status, http.StatusBadGateway)
		return
	}

	io.Copy(response, resp.Body)
}

func ServeFile(response http.ResponseWriter, request *http.Request) {
	filename := mux.Vars(request)["filename"]
	http.ServeFile(response, request, filename)
}
