package drawcard

import (
	"go-qbot/bot/api"
	"go-qbot/bot/logging"
	"go-qbot/modules/responsor"
	"io"
	"net/http"
	"net/url"
)

var logger *logging.Logs

func init() {
	logger = responsor.RegistCommand("发病", handler)
}

func getdata(name string) string {
	u := "http://seiki.fun:8008/api/diana/?name=" + url.QueryEscape(name)
	logger.Debugln(u)
	resp, err := http.Get(u)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

func handler(argv []string, gmsg api.GroupMessage) (out_msg *api.MessageChain, enable bool) {
	if len(argv) == 0 {
		return nil, false
	}
	enable = true
	msg := getdata(argv[0])
	out_msg = api.NewMessageChain().AddPlain(msg)

	return
}
