package api

import "github.com/mitchellh/mapstructure"

type MessageElement interface {
	MessageElementPlain | MessageElementSource
}

type MessageElementSource struct {
	Type string `json:"type" mapstructure:"type"`
	Id   int64  `json:"id "  mapstructure:"id"`
	Time int64  `json:"time" mapstructure:"time"`
}

type MessageElementPlain struct {
	Type string `json:"type" mapstructure:"type"`
	Text string `json:"text" mapstructure:"text"`
}

func toMap[T MessageElement](e T) map[string]interface{} {
	out := make(map[string]interface{})
	mapstructure.Decode(e, &out)
	return out
}

type MessageChain struct {
	elems []map[string]interface{}
}

func NewMessageChain() *MessageChain {
	return &MessageChain{
		elems: make([]map[string]interface{}, 0),
	}
}

func (mc *MessageChain) AddPlain(text string) *MessageChain {

	mc.elems = append(mc.elems, toMap(MessageElementPlain{Type: "Plain", Text: text}))
	return mc
}
func (mc MessageChain) ToJsonList() []map[string]interface{} {
	return mc.elems
}
