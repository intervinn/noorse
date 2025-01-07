package main

import (
	"github.com/intervinn/noorse"
	"github.com/intervinn/noorse/commands"
	"github.com/intervinn/noorse/storage"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	storage.GetInstance()
	bot := noorse.GetInstance()
	bot.Init(commands.Commands)
	bot.Serve()
}
