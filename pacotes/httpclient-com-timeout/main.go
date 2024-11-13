package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://conteudobit.com.br")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	println(string(body))
}
