package main

import (
	"log"
	"net/http"
	"time"
)

type Usuario struct {
	Nome string
}

func main() {
	http.HandleFunc("/", sendMail)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processada")
		w.Write([]byte("1234"))
	case <-ctx.Done():
		log.Println("Request cancelada")
	}
}
