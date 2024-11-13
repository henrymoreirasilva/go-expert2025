package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	sendMail(ctx)
}

func sendMail(ctx context.Context) {
	select {
	case <-ctx.Done():
		println("Cancelado por timeout")
		return
	case <-time.After(5 * time.Second):
		println("Concluído com sucesso")
	}
}
