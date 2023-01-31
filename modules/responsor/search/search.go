package drawcard

import (
	"encoding/json"
	"fmt"
	"go-qbot/bot"
	"go-qbot/bot/api"
	"go-qbot/modules/responsor"
	"io"
	"net/http"
	"net/url"
)

//var logger *logging.Logs
var token, _ = bot.GetModConf("saucenao").(string)

func init() {
	responsor.RegistCommand("search", handler)
}

type response struct {
	Results []struct {
		Header struct {
			Similarity string
			Thumbnail  string
			Index_name string
		}
		Data struct {
			Ext_urls []string
		}
	}
}

func getdata(picurl string) *api.MessageChain {
	var url = "https://saucenao.com/search.php?output_type=2&testmode=1&numres=5&url=" + url.QueryEscape(picurl) + "&api_key=" + token

	resp, err := http.Get(url)
	if err != nil {
		return api.NewMessageChain().AddPlain(err.Error())
	}
	var resstruct response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return api.NewMessageChain().AddPlain(err.Error())
	}

	err = json.Unmarshal(body, &resstruct)

	if err != nil {
		return api.NewMessageChain().AddPlain(string(body))
	}

	msg := api.NewMessageChain()
	for _, e := range resstruct.Results {
		purl := ""
		if len(e.Data.Ext_urls) > 0 {
			purl = e.Data.Ext_urls[0]
		}
		msg.
			AddImage("", e.Header.Thumbnail, "", nil).
			AddPlain(fmt.Sprintf("相似度：%v\n索引名：%v\n源：%v", e.Header.Similarity, e.Header.Index_name, purl))
	}
	return msg
}

func handler(argv []string, gmsg api.GroupMessage) (out_msg *api.MessageChain, enable bool) {
	enable = false
	if len(gmsg.MessageChain) == 3 && gmsg.MessageChain[2]["type"].(string) == "Image" {
		picurl := gmsg.MessageChain[2]["url"].(string)
		out_msg = getdata(picurl)
		if out_msg != nil {
			enable = true
		}
	}

	return
}
