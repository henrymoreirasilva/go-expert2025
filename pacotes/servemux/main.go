package main

import (
	"fmt"
	"net/http"
)

type httpResponse struct{}

func (h httpResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Acessando via Handle")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Acessando via HandleFunc")
	})
	mux.Handle("/struct", httpResponse{})

	http.ListenAndServe(":8080", mux)
}
