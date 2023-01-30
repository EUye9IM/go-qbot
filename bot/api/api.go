package api

import (
	"context"
	"go-qbot/bot/connect"
)

func SendGroupMessage(ctx context.Context, target int64, message_chain []map[string]interface{}) (map[string]interface{}, error) {
	content := make(map[string]interface{})
	content["sessionKey"] = connect.Session()
	content["target"] = target
	content["messageChain"] = message_chain
	return connect.SendCommand(
		ctx,
		"sendGroupMessage",
		"",
		content,
	)
}
