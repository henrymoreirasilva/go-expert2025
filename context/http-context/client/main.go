package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "senha", "1234")

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(compare(ctx, body))

}

func compare(ctx context.Context, body []byte) string {
	if ctx.Value("senha") == string(body) {
		return "senha correta"
	}

	return "senha incorreta"
}
