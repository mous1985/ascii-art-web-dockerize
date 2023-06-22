package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const port = ":8080"

var templates = template.Must(template.ParseFiles("Templates/stylize.html"))

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		fmt.Fprintf(w, "Status 404: Page Not Found")
		return
	}

	templates.ExecuteTemplate(w, "stylize.html", "")
}

func dynamique(w http.ResponseWriter, r *http.Request) {
	var fichier *os.File

	var s []string
	var str string
	recept := r.FormValue("fichier")

	receptxt := r.FormValue("formulaire")
	for i := range receptxt {
		if receptxt[i] < 32 || receptxt[i] > 127 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "veuillez entrer un caractere appartenant a la table ascii")
			return

		}
	}
	content := strings.ReplaceAll(receptxt, "\r\n", "\\n")

	contente := strings.Split(content, "\\n")
	if strings.Contains(receptxt, "ù") {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal Server Error"))
		log.Println(http.StatusInternalServerError)
		return

	}

	if recept == "shadow" {
		fichier, _ = os.Open("shadow.txt")
	} else if recept == "standard" {
		fichier, _ = os.Open("standard.txt")
	} else if recept == "thinkertoy" {
		fichier, _ = os.Open("thinkertoy.txt")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Status 400: Bad Request")
		return
	}

	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	for _, element := range contente {
		if len(element) > 0 {
			line := []rune(element)

			for a := 0; a < 8; a++ {

				for i := 0; i < len(line); i++ {
					group := (int(line[i]) - 32) * 9
					adress := group + a + 1

					str += (s[adress])
				}
				str += (string(rune('\n')))
			}
		} else {
			str += (string(rune('\n')))
		}
	}
	templates.ExecuteTemplate(w, "stylize.html", "")
	fmt.Fprint(w, str)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii", dynamique)

	fmt.Printf("http://localhost:8080")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("Templates/css/"))))
	http.ListenAndServe(port, nil)
}
