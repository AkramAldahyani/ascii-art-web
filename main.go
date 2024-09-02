package main
import (
	ascii "ascii/functions"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)
func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Define handlers for the main page, ASCII art result, and file export
	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii-art", resultHandler)
	http.HandleFunc("/download", downloadHandler)
	// Start the server on port 8080
	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
// index processes form values and generates ASCII art
func index(r *http.Request, w http.ResponseWriter) (string, int) {
	FILENAME := r.FormValue("temp")
	argument := r.FormValue("argu")
	if argument == "" {
		http.Error(w, "No input provided.", http.StatusBadRequest)
	}
	if FILENAME == "" || FILENAME != "standard" || FILENAME != "shadow" || FILENAME != "thinkertoy" {
		FILENAME = "standard"
	}
	if strings.Count(argument, "\\n")*2 == len(argument) {
		return "Invalid input.", http.StatusBadRequest
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
	return art, http.StatusOK
}
// handler serves the main page and redirects form submissions
func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/ascii-art", http.StatusSeeOther)
		return
	}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
// resultHandler processes form submissions and displays ASCII art
func resultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodPost {
		result, statusCode := index(r, w)
		if statusCode != http.StatusOK {
			http.Error(w, result, statusCode)
			return
		}
		// Serve the template with the ASCII art and include a download link
		tmpl := template.Must(template.ParseFiles("ascii-art.html"))
		err := tmpl.ExecuteTemplate(w, "ascii-art.html", struct {
			Art string
		}{
			Art: result,
		})
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
}
// downloadHandler handles the download request for the ASCII art
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/download" {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodPost {
		result := r.FormValue("art")
		if result == "" {
			http.Error(w, "No input provided.", http.StatusBadRequest)
			return
		}
		// Set headers to indicate a file download
		w.Header().Set("Content-Disposition", "attachment; filename=\"ascii-art.txt\"")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(result)))
		// Write the ASCII art to the response
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
}
