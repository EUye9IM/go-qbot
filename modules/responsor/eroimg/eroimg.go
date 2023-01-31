package drawcard

import (
	"encoding/json"
	"fmt"
	"go-qbot/bot/api"
	"go-qbot/bot/logging"
	"go-qbot/modules/responsor"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var logger *logging.Logs

func init() {
	rand.Seed(time.Now().Unix())
	logger = responsor.RegistCommand("ero", handler)
}

func getdata(apiurl string) string {
	apiurl += "proxy=i.pixiv.cat/{{path}}"
	u := "https://api.lolicon.app/setu/v2?" + url.QueryEscape(apiurl)
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

type res struct {
	Error string
	Data  []struct {
		Pid   int
		P     int
		Uid   int
		Title string
		Urls  struct {
			Original string
		}
	}
}

func handler(argv []string, gmsg api.GroupMessage) (out_msg *api.MessageChain, enable bool) {
	enable = true
	out_msg = api.NewMessageChain().
		At(gmsg.Sender.Id, "@"+gmsg.Sender.MemberName)

	var apiurl string
	for i := 0; i < len(argv) && i < 3; i++ {
		apiurl += "&tag=" + argv[i]
	}

	msg := getdata(apiurl)
	var r res

	err := json.Unmarshal([]byte(msg), &r)
	if err != nil {
		logger.Infof("错误：%v %v\f", msg, err.Error())
		out_msg.AddPlain(msg)
		return
	}
	if r.Error != "" {
		out_msg.AddPlain(r.Error)
		return
	}
	for _, v := range r.Data {
		out_msg.AddImage("", v.Urls.Original, "", nil).
			AddPlain(fmt.Sprintf("%v\npid_p_uid\n%v_%v_%v", v.Title, v.Pid, v.P, v.Uid))
	}

	return
}
