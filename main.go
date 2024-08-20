package main

import (
	ascii "ascii/functions"
	"html/template"
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii-art", resultHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

// This fuction return the output
func index(r *http.Request) string {
	FILENAME := r.FormValue("temp")
	argument := r.FormValue("argu")
	if argument == "" {
		return "No input provided."
	}
	if strings.Count(argument, "\\n")*2 == len(argument) {
		return "Invalid input."
	}

	var art string
	letters := ascii.Read(FILENAME)
	statements := ascii.Split(argument, "\\n")

	for _, s := range statements {
		if s == "\n" {
			art += "\n"
			continue
		}
		art += ascii.Print(s, letters)
	}
	art += "\n"

	return art
}

// Handling the main page
func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == http.MethodPost {
			// Handle form submission and redirect to results
			http.Redirect(w, r, "/ascii-art", http.StatusSeeOther)
			return
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "index.html", nil)
	} else {
		http.Error(w, "404 Page Not Found\n", 404)
	}

}

// Handling the result page
func resultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ascii-art" {
		if r.Method == http.MethodPost {
			result := index(r)

			tmpl := template.Must(template.ParseFiles("ascii-art.html"))
			tmpl.ExecuteTemplate(w, "ascii-art.html", result)
			// fmt.Fprintf(w, "<h1>Submission Result</h1>")
			// fmt.Fprintf(w, "<pre>%s</pre>", result) // Use <pre> for preserving ASCII art formatting
			return
		}
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	} else {
		http.Error(w, "404 Page Not Found\n", 404)
	}

}
