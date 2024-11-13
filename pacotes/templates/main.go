package main

import (
	"log"
	"net/http"
	"text/template"
)

type Contato struct {
	Mensagem string
}
type Sobre struct {
	Mensagem string
}

func main() {
	templateFiles := []string{
		"header.html",
		"sobre.html",
		"contato.html",
		"footer.html",
	}

	http.HandleFunc("/sobre", func(w http.ResponseWriter, r *http.Request) {
		templates := template.Must(template.New("sobre.html").ParseFiles(templateFiles...))
		data := Sobre{Mensagem: "Conheça nossa empresa"}

		err := templates.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}

	})

	http.HandleFunc("/contato", func(w http.ResponseWriter, r *http.Request) {
		templates := template.Must(template.New("contato.html").ParseFiles(templateFiles...))
		data := Contato{Mensagem: "Entre em contato"}

		err := templates.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}

	})
	http.ListenAndServe(":8080", nil)
}
