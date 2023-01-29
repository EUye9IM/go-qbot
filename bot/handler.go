package bot

import "go-qbot/bot/api"

type Handler[T api.Message] func(T)

var (
	handlers []interface{} = make([]interface{}, 0)
)

func RegistHandlers[T api.Message](h Handler[T]) {
	handlers = append(handlers, h)
}
