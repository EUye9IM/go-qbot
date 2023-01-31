package main

import (
	"go-qbot/bot"

	_ "go-qbot/modules/responsor"
	_ "go-qbot/modules/responsor/drawcard"
	_ "go-qbot/modules/responsor/eroimg"
)

func main() {
	bot.Run()
}
