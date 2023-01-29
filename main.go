package main

import (
	"fmt"
	"go-qbot/bot"
	"go-qbot/config"

	_ "go-qbot/modules/drawcard"
)

func main() {
	//log
	fmt.Printf("Config: %+v\n", config.Conf())
	bot.Run()
}
