package connect

import (
	"context"
	"fmt"
	"go-qbot/config"
	"net/http"
	"strconv"

	"github.com/mitchellh/mapstructure"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type response struct {
	SyncId string      `json:"syncId"`
	Data   interface{} `json:"data"`
}

type request struct {
	SyncId     int         `json:"syncId"`
	Command    string      `json:"command"`
	SubCommand string      `json:"subCommand"`
	Content    interface{} `json:"content"`
}

var (
	ctx     context.Context
	conn    *websocket.Conn
	session string
	chmap   map[int]chan interface{} = make(map[int]chan interface{})
	next_id int                      = 1
)

func init() {
	ctx = context.Background()
}

func Connect(cfg_conn config.ConfigConnect) error {
	var err error
	header := &http.Header{}
	if cfg_conn.Qq != "" {
		header.Add("qq", cfg_conn.Qq)
	}
	if cfg_conn.Verify_key != "" {
		header.Add("verifyKey", cfg_conn.Verify_key)
	}
	conn, _, err = websocket.Dial(
		ctx,
		fmt.Sprintf("ws://%v:%v/all", cfg_conn.Host, cfg_conn.Port),
		&websocket.DialOptions{HTTPHeader: *header})
	if err != nil {
		return fmt.Errorf("connect failed: %w", err)
	}
	return nil
}
func Disconnect(reason string) error {
	err := conn.Close(websocket.StatusNormalClosure, reason)
	if err != nil {
		return fmt.Errorf("disconnect failed: %w", err)
	}
	return nil
}

func read() (*response, error) {
	res := &response{}
	err := wsjson.Read(ctx, conn, &res)
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}
	return res, nil
}

func write(req request) error {
	err := wsjson.Write(ctx, conn, req)
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	return nil
}

func Session() string {
	return session
}

func RecvData() (interface{}, error) {
	for {
		res, err := read()
		if err != nil {
			return nil, fmt.Errorf("recv data failed: %w", err)
		}
		if res.SyncId == "" {
			// session
			var conn_data struct {
				Code    float64
				Session string
			}
			err := mapstructure.Decode(res.Data, &conn_data)
			if err != nil {
				return res.Data, fmt.Errorf("recv data failed: unknown data: %w:{{ %#v }}", err, res)
			}
			session = conn_data.Session
		} else {
			syncid, _ := strconv.Atoi(res.SyncId)
			if syncid < 0 {
				return res.Data, nil
			} else {
				ch, ok := chmap[syncid]
				if ok {
					ch <- res.Data
				} else {
					return res.Data, fmt.Errorf("recv data failed: unknown syncid: %+v", res)
				}
			}
		}
	}
}

func registChannel() (int, error) {
	const MAX_SYNCID = 1000
	for i := 0; i < MAX_SYNCID; i++ {
		id := (next_id+i+MAX_SYNCID-1)%MAX_SYNCID + 1
		_, ok := chmap[id]
		if !ok {
			chmap[id] = make(chan interface{})
			return id, nil
		}
	}
	return 0, fmt.Errorf("regist channel failed: please increase MAX_SYNCID")
}

func SendCommand(ctx context.Context, command, subcommand string, content interface{}) (interface{}, error) {
	syncid, err := registChannel()
	if err != nil {
		return nil, fmt.Errorf("send commant failed: %w", err)
	}

	req := request{
		SyncId:     syncid,
		Command:    command,
		SubCommand: subcommand,
		Content:    content,
	}
	err = write(req)
	if err != nil {
		return nil, fmt.Errorf("send commant failed: %w", err)
	}

	select {
	case data := <-chmap[syncid]:
		delete(chmap, syncid)
		return data, nil
	case <-ctx.Done():
	}
	return nil, fmt.Errorf("send commant failed: %w", ctx.Err())
}
