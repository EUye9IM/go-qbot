package api

import (
	"encoding/base64"

	"github.com/mitchellh/mapstructure"
)

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
func (mc *MessageChain) AddQuote(
	id int, groupid int64, senderid int64,
	orig_msgchain []map[string]interface{},
) *MessageChain {
	targetid := groupid
	if targetid == 0 {
		targetid = senderid
	}
	mc.elems = append(mc.elems, toMap(
		MessageElementQuote{
			Type:     "Quote",
			Id:       id,
			GroupId:  groupid,
			SenderId: senderid,
			TargetId: targetid,
			Origin:   orig_msgchain,
		}))
	return mc
}
func (mc *MessageChain) At(target int64, display string) *MessageChain {
	mc.elems = append(mc.elems, toMap(
		MessageElementAt{
			Type:    "At",
			Target:  target,
			Display: display,
		}))
	return mc
}
func (mc *MessageChain) AddPlain(text string) *MessageChain {
	mc.elems = append(mc.elems, toMap(MessageElementPlain{Type: "Plain", Text: text}))
	return mc
}
func (mc *MessageChain) AddImage(imageId, url, path string, raw []byte) *MessageChain {
	b64 := ""
	if len(raw) > 0 {
		b64 = base64.StdEncoding.EncodeToString(raw)
	}
	mc.elems = append(mc.elems, toMap(MessageElementImage{
		Type:    "Image",
		ImageId: imageId,
		Url:     url,
		Path:    path,
		Base64:  b64,
	}))
	return mc
}
func (mc MessageChain) ToJsonList() []map[string]interface{} {
	return mc.elems
}
