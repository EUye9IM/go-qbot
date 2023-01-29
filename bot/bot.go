package bot

import (
	"fmt"
	"go-qbot/bot/api"
	"go-qbot/bot/connect"
	"go-qbot/config"
	"time"

	"github.com/mitchellh/mapstructure"
)

func classify(data map[string]interface{}) error {
	wrongdataerr := fmt.Errorf("wrong data format: %+v", data)
	datatype, ok := data["type"].(string)
	if !ok {
		return wrongdataerr
	}
	switch datatype {
	case "GroupMessage":
		var msg api.GroupMessage
		err := mapstructure.Decode(data, &msg)
		if err != nil {
			return wrongdataerr
		}
		for _, h := range handlers {
			hfunc := h.(Handler[api.GroupMessage])
			if hfunc != nil {
				go hfunc(msg)
			}
		}
	default:
		return fmt.Errorf("unknown data format: %+v", data)
	}
	return nil
}

func loop() {
	for {
		data, err := connect.RecvData()
		if err != nil {
			//log
			fmt.Println(err)
		} else {
			classify(data)
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
