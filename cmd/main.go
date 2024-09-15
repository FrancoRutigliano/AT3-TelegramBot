package main

import (
	"atomico3bot/internal/bot"
	"log"
)

func main() {
	if err := bot.StartBot(); err != nil {
		log.Fatal("error al inicializar el bot")
	}

	log.Println("Connected")
}
