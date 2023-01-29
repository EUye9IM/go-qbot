package api

type MessageElementSource struct {
	Type string `json:"type"`
	Id   int64  `json:"id"`
	Time int64  `json:"time"`
}

type MessageElementPlain struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type MessageChain struct {
	elems []interface{}
}

func NewMessageChain() *MessageChain {
	return &MessageChain{
		elems: make([]interface{}, 0),
	}
}

func (mc *MessageChain) AddPlain(text string) *MessageChain {
	mc.elems = append(mc.elems, MessageElementPlain{Type: "Plain", Text: text})
	return mc
}
func (mc MessageChain) End() []interface{} {
	return mc.elems
}
