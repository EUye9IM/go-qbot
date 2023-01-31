package bot

import (
	"fmt"
	"go-qbot/bot/api"
	"go-qbot/bot/connect"
	"go-qbot/bot/logging"
	"go-qbot/config"
	"time"

	"github.com/mitchellh/mapstructure"
)

var logger = logging.New("Bot")

func Modlog(name string) *logging.Logs {
	return logging.New("Modules: " + name)
}

func classify(data map[string]interface{}) error {
	wrongdataerr := fmt.Errorf("错误 data 格式: %+v", data)
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
		return fmt.Errorf("未知 data 格式: %+v", data)
	}
	return nil
}

func loop() {
	for {
		data, err := connect.RecvData()
		if err != nil {
			logger.Infoln(err)
			return
		} else {
			err := classify(data)
			if err != nil {
				logger.Infoln(err)
			}
		}
	}
}

func Run() {
	conf := config.Conf()
	logger.Debugf("bot config: %#v\n", conf)
	for retrycnt := 0; retrycnt <= int(conf.Connect.Max_retry); retrycnt++ {
		err := connect.Connect(conf.Connect)
		if err == nil {
			logger.Infoln("连接成功")
			loop()
			connect.Disconnect("end loop")
			logger.Infoln("断开连接")
		} else {
			logger.Infoln(err)
		}
		logger.Infof("%v 秒后尝试重新连接\n", conf.Connect.Retry_seconds)
		time.Sleep(time.Duration(conf.Connect.Retry_seconds) * time.Second)
	}
}
