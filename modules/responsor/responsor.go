package responsor

import (
	"context"
	"go-qbot/bot"
	"go-qbot/bot/api"
	"go-qbot/bot/logging"
	"strings"
	"sync"
	"time"

	"github.com/mitchellh/mapstructure"
)

var (
	logger     = bot.Modlog("Responsor")
	prefixes   = [2]string{".", "。"}
	responsors = make(map[string]Responsor, 0)
	mutex      sync.Mutex
)

const TIMEOUT = time.Duration(20) * time.Second

type Responsor func(argv []string, gmsg api.GroupMessage) (out_msg *api.MessageChain, enable bool)

func RegistCommand(command string, h Responsor) *logging.Logs {
	mutex.Lock()
	defer mutex.Unlock()
	_, ok := responsors[command]
	if ok {
		logger.Infoln("Error: 重复命令：" + command)
		return nil
	}
	responsors[command] = h
	logger.Debugln("载入命令 " + command)
	return logger.New(command)
}

func handler(gmsg api.GroupMessage) {
	if len(gmsg.MessageChain) >= 2 {
		var msg api.MessageElementPlain
		if mapstructure.Decode(gmsg.MessageChain[1], &msg) == nil {
			for _, pref := range prefixes {
				if strings.HasPrefix(msg.Text, pref) {
					commandline := strings.Fields(strings.TrimPrefix(msg.Text, pref))
					if len(commandline) > 0 {
						r := responsors[commandline[0]]
						if r != nil {
							msg_chain, ena := r(commandline[1:], gmsg)
							if ena {
								ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
								defer cancel()
								_, err :=
									api.SendGroupMessage(
										ctx,
										gmsg.Sender.Group.Id,
										msg_chain.ToJsonList(),
									)
								if err != nil {
									logger.Debugln(err)
								}
							}
						}
					}

					return
				}
			}
		}
	}
}

func init() {
	bot.RegistHandlers(handler)
}
