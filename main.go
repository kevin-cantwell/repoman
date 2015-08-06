package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kevin-cantwell/repoman/templates"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc(`/{filename:.*\.md}`, ServeGFM).Methods("GET")
	r.HandleFunc(`/{filename:.*}`, ServeFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":9999", r))
}

func ServeFile(response http.ResponseWriter, request *http.Request) {
	filename := mux.Vars(request)["filename"]
	if filename == "" {
		filename = "."
	}

	f, err := os.Open(filename)
	if err != nil {
		http.Error(response, "Failed to read file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(response, "Failed to read file stats: "+err.Error(), http.StatusInternalServerError)
		return
	}

	page := templates.GithubPage{
		Files: []templates.File{},
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		// TODO: Render regular files with github styling
		http.ServeFile(response, request, filename)
		return
	case mode.IsDir():
		fis, err := ioutil.ReadDir(filename)
		if err != nil {
			http.Error(response, "Failed to read directory: "+err.Error(), http.StatusInternalServerError)
			return
		}
		for _, fi := range fis {
			file := templates.File{
				Name:  fi.Name(),
				Path:  filename + string(os.PathSeparator) + fi.Name(),
				IsDir: fi.IsDir(),
			}
			if strings.ToLower(fi.Name()) == "readme.md" {
				gfmBody, err := ReadAsGFM(file.Path)
				if err != nil {
					http.Error(response, "Failed to parse README as Github-flavored markdown: "+err.Error(), http.StatusInternalServerError)
					return
				}
				page.GFM = template.HTML(gfmBody)
			}
			page.Files = append(page.Files, file)
		}
	}

	t, err := template.New("github").Parse(templates.GithubTemplate)
	if err != nil {
		http.Error(response, "Error parsing Github-flavored markdown template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(response, page)
}

func ServeGFM(response http.ResponseWriter, request *http.Request) {
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

	markdownBody, err := ioutil.ReadAll(resp.Body)
	page := templates.GithubPage{
		GFM: template.HTML(markdownBody),
	}

	t, err := template.New("gfm").Parse(templates.GithubTemplate)
	if err != nil {
		http.Error(response, "Error parsing Github-flavored markdown template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(response, page)
}

func ReadAsGFM(filename string) ([]byte, error) {
	markdown, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	body := map[string]string{
		"text": string(markdown),
		"mode": "gfm",
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	r, _ := http.NewRequest("POST", "https://api.github.com/markdown", bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != 200 {
		return nil, errors.New("error: unexpected response from github api: " + resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
