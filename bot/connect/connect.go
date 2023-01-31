package connect

import (
	"context"
	"encoding/json"
	"fmt"
	"go-qbot/bot/logging"
	"go-qbot/config"
	"net/http"
	"strconv"
	"sync"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type response struct {
	SyncId string                 `json:"syncId"`
	Data   map[string]interface{} `json:"data"`
}

type request struct {
	SyncId     int         `json:"syncId"`
	Command    string      `json:"command"`
	SubCommand string      `json:"subCommand"`
	Content    interface{} `json:"content"`
}

var (
	ctx      context.Context
	conn     *websocket.Conn
	session  string
	mapmutex sync.RWMutex
	chmap    map[int]chan map[string]interface{} = make(map[int]chan map[string]interface{})
	next_id  int                                 = 1

	logger = logging.New("Connect")
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
	logger.Debugf("read：%+v\n", *res)
	return res, nil
}

func write(req request) error {
	err := wsjson.Write(ctx, conn, req)
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	b, _ := json.Marshal(req)
	logger.Debugf("write：%v\n", string(b))
	return nil
}

func Session() string {
	return session
}

func RecvData() (map[string]interface{}, error) {
	for {
		res, err := read()
		if err != nil {
			return nil, fmt.Errorf("recv data failed: %w", err)
		}
		if res.SyncId == "" {
			// session
			s, ok := res.Data["session"]
			if ok {
				ss, ok2 := s.(string)
				if ok2 {
					session = ss
				} else {
					ok = false
				}
			}
			if !ok {
				return nil, fmt.Errorf("recv data failed: unknown data: %#v", res)
			}
		} else {
			syncid, _ := strconv.Atoi(res.SyncId)
			if syncid < 0 {
				return res.Data, nil
			} else {
				mapmutex.RLock()
				ch, ok := chmap[syncid]
				mapmutex.RUnlock()
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
	mapmutex.Lock()
	defer mapmutex.Unlock()
	const MAX_SYNCID = 1000
	for i := 0; i < MAX_SYNCID; i++ {
		id := (next_id+i+MAX_SYNCID-1)%MAX_SYNCID + 1
		_, ok := chmap[id]
		if !ok {
			chmap[id] = make(chan map[string]interface{})
			next_id = id + 1
			return id, nil
		}
	}
	return 0, fmt.Errorf("regist channel failed: please increase MAX_SYNCID")
}
func cancelChannel(syncid int) {
	mapmutex.Lock()
	defer mapmutex.Unlock()
	delete(chmap, syncid)
}

func SendCommand(ctx context.Context, command, subcommand string, content interface{}) (map[string]interface{}, error) {
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

	mapmutex.RLock()
	ch := chmap[syncid]
	mapmutex.RUnlock()
	select {
	case data := <-ch:
		cancelChannel(syncid)
		return data, nil
	case <-ctx.Done():
	}
	return nil, fmt.Errorf("send commant failed: %w", ctx.Err())
}
