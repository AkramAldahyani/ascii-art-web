package main

import (
	ascii "ascii/functions"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
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

func index() string {
	if len(os.Args) != 4 {
		fmt.Printf("Error: expected 3 argument but recieved %d\n", len(os.Args)-1)
		os.Exit(1)
	}
	FILENAME := os.Args[3]
	templates := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}
	if !templates[FILENAME] {
		fmt.Println("Error: Invalid Bannar, Please choose standard, shadow, template or thinkertoy")
		os.Exit(1)
	}
	argument := os.Args[2]
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
	useColor := flag.String("color", "Reset", "Choose one color: Reset, Black, Red, Green, Yellow or Blue")
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

type Art struct {
	art string
}

func main() {
	hello := index()
	hi := Art{art: hello}
	fmt.Println(hi.art)
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
}

var tpl *template.Template

func handler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "index.html", nil)
}
func init() {
	tpl = template.Must(template.ParseGlob("index.html"))
}
