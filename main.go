package main

import (
	"go-qbot/bot"

	_ "go-qbot/modules/responsor"
	_ "go-qbot/modules/responsor/drawcard"
	_ "go-qbot/modules/responsor/eroimg"
	_ "go-qbot/modules/responsor/search"
	_ "go-qbot/modules/responsor/sick"
)

func main() {
	bot.Run()
}
