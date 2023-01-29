package main

import (
	"context"
	"fmt"
	"go-qbot/bot"
	"go-qbot/bot/connect"
	"go-qbot/config"
	"time"
)

func h_draw(data interface{}) {
	// rubbish
	datajson, ok := data.(map[string]interface{})
	if !ok {
		return
	}
	datatype, ok := datajson["type"].(string)
	if !ok {
		return
	}
	if datatype == "GroupMessage" {
		msgc, ok := datajson["messageChain"].([]interface{})
		if !ok {
			return
		}
		for _, onemsg := range msgc {
			onemsgmap, ok := onemsg.(map[string]interface{})
			if !ok {
				break
			}
			if onemsgmap["type"].(string) == "Plain" {
				if onemsgmap["text"].(string) == "。draw 单张塔罗牌" {
					sender := datajson["sender"].(map[string]interface{})
					name := sender["memberName"].(string)
					target := sender["group"].(map[string]interface{})["id"].(float64)
					msg := draw_card(name)

					content := make(map[string]interface{})

					content["sessionKey"] = connect.Session()
					content["target"] = target
					msgs := make([]map[string]interface{}, 1)
					msgs[0] = make(map[string]interface{})
					msgs[0]["type"] = "Plain"
					msgs[0]["text"] = msg

					content["messageChain"] = msgs
					ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
					rsp, err := connect.SendCommand(ctx, "sendGroupMessage", "", content)
					fmt.Println(rsp, err)
				}
				break
			}
		}
	}
}

func main() {
	//log
	fmt.Printf("Config: %+v\n", config.Conf())
	bot.RegistHandlers(h_draw)
	bot.Run()
}
