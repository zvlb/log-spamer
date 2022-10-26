package main

import (
	"context"
	"log"

	"github.com/zvlb/log-spammer/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Это настройка, чтобы программа останавливалась через 5 мекунд. Удобно для Дебага.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5000)
	// defer cancel()

	if err := app.NewApp(ctx); err != nil {
		log.Fatalln(err)
	}
}
