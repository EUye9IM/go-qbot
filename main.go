package main

import (
	"go-qbot/bot"

	_ "go-qbot/modules/responsor"
	_ "go-qbot/modules/responsor/drawcard"
)

func main() {
	bot.Run()
}
