package main

import (
	"bufio"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("index.html"))

	style := http.FileServer(http.Dir("./fonts"))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", style))
	http.HandleFunc("/ascii-art", posthandler)
	http.HandleFunc("/", gethandler)

	go openBrowser("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func openBrowser(url string) {
	err := open.Run(url)
	if err != nil {
		panic(err)
	}
}

func gethandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path != "/" {
			file, err := os.Open("error.html")
			if err != nil {
				http.Error(w, "An error occurred.", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// Set the Content-Type header
			w.Header().Set("Content-Type", "text/html")

			// Write the contents of the error.html file to the response writer
			if _, err := io.Copy(w, file); err != nil {
				http.Error(w, "An error occurred.", http.StatusInternalServerError)
				return
			}
			return
		}
		renderTemplate(w, "index.html", nil)
	}
}

func posthandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		text1 := r.FormValue("text")
		font := r.FormValue("fonts")
		text := cleanText(text1)

		if !isValidASCII(text) {
			http.Error(w, "ERROR-400\nBad request!", http.StatusBadRequest)
			return
		}

		file, err := os.Open("fonts/" + font + ".txt")
		if err != nil {
			http.Error(w, "ERROR-400\nBad Request!! \nPlease make sure you select a font.", http.StatusBadRequest)
			return
		}
		defer file.Close()

		lines, err := readLines(file)
		if err != nil {
			http.Error(w, "ERROR-500\nInternal Server Error", http.StatusInternalServerError)
			return
		}

		asciiChrs := identifyASCIIChars(lines)

		var c string
		for i := 0; i < len(text); i++ {
			if text[i] == 92 && text[i+1] == 110 {
				c = PrintArt(text[:i], asciiChrs) + PrintArt(text[i+2:], asciiChrs)
			}
		}

		if !strings.Contains(text, "\n") {
			c = PrintArt(text, asciiChrs)
		}

		renderTemplate(w, "index.html", c)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, "ERROR-500\nInternal Server Error", http.StatusInternalServerError)
		return
	}
}

func cleanText(text1 string) string {
	if strings.Contains(text1, "\r\n") {
		return strings.ReplaceAll(text1, "\r\n", "\\n")
	}
	return text1
}

func isValidASCII(text string) bool {
	for _, v := range text {
		if !(v >= 32 && v <= 126) {
			return false
		}
	}
	return true
}

func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func identifyASCIIChars(lines []string) map[int][]string {
	asciiChrs := make(map[int][]string)
	dec := 31
	for _, line := range lines {
		if line == "" {
			dec++
		} else {
			asciiChrs[dec] = append(asciiChrs[dec], line)
		}
	}
	return asciiChrs
}

func PrintArt(n string, y map[int][]string) string {
	a := []string{}
	for j := 0; j < len(y[32]); j++ {
		for _, letter := range n {
			a = append(a, y[int(letter)][j])
		}
		a = append(a, "\n")
	}
	return strings.Join(a, "")
}
