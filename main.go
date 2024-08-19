package main

import (
	ascii "ascii/functions"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var clrs = map[string]string{
	"Black":  "\u001b[30m",
	"Red":    "\u001b[31m",
	"Green":  "\u001b[32m",
	"Yellow": "\u001b[33m",
	"Blue":   "\u001b[34m",
	"Reset":  "\u001b[0m",
}

func index(w http.ResponseWriter, r *http.Request) string {

	FILENAME := r.FormValue("temp")

	argument := r.FormValue("argu")
	// if argument is empty
	if argument == "" {
		return ""
	}
	// if argument is only new lines
	if strings.Count(argument, "\\n")*2 == len(argument) {
		for i := 0; i < len(argument)/2; i++ {
			fmt.Println()
		}
		return ""
	}
	var art string
	letters := ascii.Read(FILENAME)
	// Split the argument based on new line
	statments := ascii.Split(argument, "\\n")
	useColor := flag.String("color", "Red", "Choose one color: Reset, Black, Red, Green, Yellow or Blue")
	flag.Parse()
	if color, exists := clrs[*useColor]; exists {
		for _, s := range statments {
			if s == "\n" {
				fmt.Println()
				continue
			}
			art += color
			art += ascii.Print(s, letters)
		}
		art += "\n"

	} else {
		// Print an error message if the color doesn't exist
		fmt.Println("Color not recognized Please choose one from the following: Red, Black, Green, Yellow or Blue")
	}
	return art

}


func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii-art", resultHandler)
	http.ListenAndServe("", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		head string
	}{
		head: "Hello world",
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, fmt.Sprintf("/ascii-art%s", index(w,r)), http.StatusSeeOther)

}
func resultHandler(w http.ResponseWriter, r *http.Request) {
    // Retrieve query parameters
   

    // Display the result
    fmt.Fprintf(w, "<h1>Submission Result</h1>")
    fmt.Fprintf(w, "<p> %s</p>", index(w,r))
   
}