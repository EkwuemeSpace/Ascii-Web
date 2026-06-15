package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Data struct {
	Message string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "status not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	var result strings.Builder

	if r.URL.Path != "/ascii-art" {
		http.Error(w, "status not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPost {
		user := r.FormValue("text")
		banner := r.FormValue("banner")

		if user == "" || banner == "" {
			http.Error(w, "input or banner cannot be empty", http.StatusBadRequest)
			return
		}
		content, err := readBannerFile(banner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		charMap, err := parseBannerFile(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = renderString(&result, charMap, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := Data{
			Message: result.String(),
		}

		temp, err := template.ParseFiles("template/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := temp.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
