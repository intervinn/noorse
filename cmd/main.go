package main

import (
	"github.com/intervinn/noorse"
	"github.com/intervinn/noorse/commands"
	"github.com/intervinn/noorse/storage"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	storage.Instance()
	bot := noorse.Instance()
	bot.Init(commands.Commands)
	commands.Init()
	bot.Serve()
}
