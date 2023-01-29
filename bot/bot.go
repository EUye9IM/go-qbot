package bot

import (
	"fmt"
	"go-qbot/bot/connect"
	"go-qbot/config"
	"time"
)

func loop() {
	for {
		data, err := connect.RecvData()
		if err != nil {
			//log
			fmt.Println(err)
		} else {
			for _, h := range handlers {
				go h(data)
			}
		}
	}
}

func Run() {
	conf := config.Conf()
	for retrycnt := 0; retrycnt <= int(conf.Connect.Max_retry); retrycnt++ {
		err := connect.Connect(conf.Connect)
		if err != nil {
			//log
			fmt.Println(err)
			time.Sleep(time.Duration(conf.Connect.Retry_seconds) * time.Second)
			continue
		}
		//log
		fmt.Println("connect success")
		//connected
		loop()
	}
}
