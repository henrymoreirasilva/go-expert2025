package main

import "net/http"

func main() {
	// Attaches helloWord funcrtion to rout "\"
	http.HandleFunc("/", helloWorld)

	// Listening on localhost:8080
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Configure http status code
		w.WriteHeader(http.StatusForbidden)
		return
	}

	p := r.URL.Query().Get("h")
	if p == "" {
		// Configure http status code
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Configure http status code
	w.WriteHeader(http.StatusFound)

	// Configure http content type
	w.Header().Set("Condtent-type", "text/plain")

	// Write response text
	w.Write([]byte("Hello World"))
}
