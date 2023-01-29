package bot

type Handler func(interface{})

var handlers []Handler = make([]Handler, 0)

func RegistHandlers(h Handler) {
	handlers = append(handlers, h)
}
