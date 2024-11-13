package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	dados := map[string]string{
		"nome":  "henry",
		"email": "henry@zoomwi.com.br",
	}

	json, err := json.Marshal(dados)
	if err != nil {
		log.Fatal(err)
	}

	requestBody := bytes.NewBuffer(json)

	resp, err := http.Post("https://conteudobit.com.br", "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	println(string(resp.Status))
}
