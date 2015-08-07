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
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/kevin-cantwell/repoman"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/assets/github/index.css", ServeGithubCSS).Methods("GET")
	r.HandleFunc("/assets/github2/index.css", ServeGithub2CSS).Methods("GET")
	r.HandleFunc("/assets/octicons/octicons/octicons.woff", ServeFontFaceWOFF).Methods("GET")
	r.HandleFunc("/assets/octicons/octicons/octicons.eot", ServeFontFaceEOT).Methods("GET")
	r.HandleFunc("/assets/octicons/octicons/octicons.ttf", ServeFontFaceTTF).Methods("GET")
	r.HandleFunc("/assets/octicons/octicons/octicons.svg", ServeFontFaceSVG).Methods("GET")
	r.HandleFunc(`/{filename:.*}`, ServeStyledFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":9999", r))
}

var (
	wd, _ = os.Getwd()
)

func ServeGithubCSS(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/css")
	response.Write(repoman.GithubCSS)
}

func ServeGithub2CSS(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/css")
	response.Write(repoman.Github2CSS)
}

func ServeFontFaceWOFF(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/octet-stream")
	response.Write(repoman.OcticonsWOFF)
}

func ServeFontFaceEOT(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/octet-stream")
	response.Write(repoman.OcticonsEOT)
}

func ServeFontFaceTTF(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/octet-stream")
	response.Write(repoman.OcticonsTTF)
}

func ServeFontFaceSVG(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/octet-stream")
	response.Write(repoman.OcticonsSVG)
}

func ServeStyledFile(response http.ResponseWriter, request *http.Request) {
	filename := mux.Vars(request)["filename"]

	fileToOpen := filename
	if filename == "" {
		fileToOpen = wd
	}

	f, err := os.Open(fileToOpen)
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

	baseFilename := filepath.Base(filename)
	if baseFilename == "." {
		baseFilename = ""
	}

	page := repoman.GithubPage{
		RepoName:     filepath.Base(wd),
		BaseFilename: baseFilename,
		IsDir:        fi.IsDir(),
		Breadcrumbs:  Breadcrumbs(filename),
		Files:        []repoman.File{},
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		if !strings.HasSuffix(fi.Name(), ".md") {
			// If it's not a markdown file...
			http.ServeFile(response, request, fileToOpen)
			return
		} else {
			// If it is a markdown file...
			gfmBody, err := ReadAsGFM(fileToOpen)
			if err != nil {
				http.Error(response, "Failed to parse README as Github-flavored markdown: "+err.Error(), http.StatusInternalServerError)
				return
			}
			page.GFM = template.HTML(gfmBody)
		}
	case mode.IsDir():
		subfis, err := ioutil.ReadDir(fileToOpen)
		if err != nil {
			http.Error(response, "Failed to read directory: "+err.Error(), http.StatusInternalServerError)
			return
		}
		for _, subfi := range subfis {
			path := "/" + subfi.Name()
			if filename != "" {
				path = "/" + filename + path
			}
			file := repoman.File{
				Name:  subfi.Name(),
				Path:  path,
				IsDir: subfi.IsDir(),
			}
			if strings.ToLower(subfi.Name()) == "readme.md" {
				gfmBody, err := ReadAsGFM(fileToOpen + string(os.PathSeparator) + subfi.Name())
				if err != nil {
					http.Error(response, "Failed to parse README as Github-flavored markdown: "+err.Error(), http.StatusInternalServerError)
					return
				}
				page.GFM = template.HTML(gfmBody)
			}
			page.Files = append(page.Files, file)
		}
	}

	t, err := template.New("github").Parse(repoman.GithubTemplate)
	if err != nil {
		http.Error(response, "Error parsing Github-flavored markdown template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(response, page)
}

func Breadcrumbs(filename string) []repoman.Breadcrumb {
	var path string
	crumbs := []repoman.Breadcrumb{}
	components := strings.Split(filename, "/")
	for i, component := range components {
		// Breadcrumbs never include the file's name itself
		if i == len(components)-1 {
			break
		}
		path += "/" + component
		crumb := repoman.Breadcrumb{
			Path: path,
			Name: component,
		}
		crumbs = append(crumbs, crumb)
	}
	return crumbs
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

	log.Printf("Posting %v to Github's API for markdown processing...\n", filename)
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

	reset, _ := strconv.ParseInt(resp.Header.Get("X-RateLimit-Reset"), 10, 64)
	log.Printf(
		"Github API rate limits: limit=%v, remaining=%v, reset=%v",
		resp.Header.Get("X-RateLimit-Limit"),
		resp.Header.Get("X-RateLimit-Remaining"),
		(time.Unix(reset, 0).Sub(time.Now())/time.Second)*time.Second, // Truncate sub-second values.
	)

	if resp.StatusCode != 200 {
		return nil, errors.New("error: unexpected response from github api: " + resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
