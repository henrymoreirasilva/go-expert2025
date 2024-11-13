package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	request, err := http.NewRequest("GET", "https://conteudobit.com.br", nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", "text/html")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	println(string(body))

}
