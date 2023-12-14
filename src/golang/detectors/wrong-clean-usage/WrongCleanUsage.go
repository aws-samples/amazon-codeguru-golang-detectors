package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const root = "/tmp"

// {fact rule=path-traversal@v1.0 defects=1}
func filepathCleanMisuseNoncompliant() {
	mux := http.NewServeMux()
	mux.HandleFunc("/myfiles", func(w http.ResponseWriter, r *http.Request) {
		filename := filepath.FromSlash(filepath.Join(root, strings.Trim(filepath.Clean(r.URL.Path), "/")))
		// Noncompliant: filepath clean is used.
		contents, err := os.ReadFile(filename)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Write(contents)
	})
}

// {/fact}

// {fact rule=path-traversal@v1.0 defects=0}
func filepathCleanMisuseCompliant() {
	mux := http.NewServeMux()
	mux.HandleFunc("/myfiles", func(w http.ResponseWriter, r *http.Request) {
		filename := filepath.FromSlash(filepath.Join(root, strings.Trim(r.URL.Path, "/")))
		// Compliant: filepath clean not used.
		contents, err := os.ReadFile(filename)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Write(contents)
	})
}

// {/fact}
