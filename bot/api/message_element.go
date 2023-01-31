package api

type MessageElement interface {
	MessageElementSource |
		MessageElementQuote |
		MessageElementAt |
		MessageElementAtAll |
		MessageElementFace |
		MessageElementPlain |
		MessageElementImage |
		MessageElementFlashImage |
		MessageElementVoice |
		MessageElementXml |
		MessageElementJson |
		MessageElementApp |
		MessageElementPoke |
		MessageElementDice |
		MessageElementMarkectFace |
		MessageElementMusicShare |
		MessageElementForwardMessage |
		MessageElementFile
}

type MessageElementSource struct {
	Type string `json:"type" mapstructure:"type"`
	Id   int64  `json:"id"   mapstructure:"id"`
	Time int64  `json:"time" mapstructure:"time"`
}
type MessageElementQuote struct {
	Type     string                   `json:"type"     mapstructure:"type"`
	Id       int                      `json:"id"       mapstructure:"id"`
	GroupId  int64                    `json:"groupId"  mapstructure:"groupId"`
	SenderId int64                    `json:"senderId" mapstructure:"senderId"`
	TargetId int64                    `json:"targetId" mapstructure:"targetId"`
	Origin   []map[string]interface{} `json:"origin"   mapstructure:"orgin"`
}
type MessageElementAt struct {
	Type    string `json:"type"    mapstructure:"type"`
	Target  int64  `json:"target"  mapstructure:"target"`
	Display string `json:"display" mapstructure:"display"`
}
type MessageElementAtAll struct {
	Type string `json:"type" mapstructure:"type"`
}
type MessageElementFace struct {
	Type   string `json:"type"             mapstructure:"type"`
	FaceId string `json:"faceId,omitempty" mapstructure:"faceId,omitempty"`
	Name   string `json:"name,omitempty"   mapstructure:"name,omitempty"`
}
type MessageElementPlain struct {
	Type string `json:"type" mapstructure:"type"`
	Text string `json:"text" mapstructure:"text"`
}
type MessageElementImage struct {
	Type    string `json:"type"              mapstructure:"type"`
	ImageId string `json:"imageId,omitempty" mapstructure:"imageId,omitempty"`
	Url     string `json:"url,omitempty"     mapstructure:"url,omitempty"`
	Path    string `json:"path,omitempty"    mapstructure:"path,omitempty"`
	Base64  string `json:"base64,omitempty"  mapstructure:"base64,omitempty"`
}
type MessageElementFlashImage struct {
	Type    string `json:"type"              mapstructure:"type"`
	ImageId string `json:"imageId,omitempty" mapstructure:"imageId,omitempty"`
	Url     string `json:"url,omitempty"     mapstructure:"url,omitempty"`
	Path    string `json:"path,omitempty"    mapstructure:"path,omitempty"`
	Base64  string `json:"base64,omitempty"  mapstructure:"base64,omitempty"`
}
type MessageElementVoice struct {
	Type    string `json:"type"              mapstructure:"type"`
	VoiceId string `json:"voiceId,omitempty" mapstructure:"voiceId,omitempty"`
	Url     string `json:"url,omitempty"     mapstructure:"url,omitempty"`
	Path    string `json:"path,omitempty"    mapstructure:"path,omitempty"`
	Base64  string `json:"base64,omitempty"  mapstructure:"base64,omitempty"`
	Length  string `json:"length,omitempty"  mapstructure:"length,omitempty"`
}
type MessageElementXml struct {
	Type string `json:"type" mapstructure:"type"`
	json string `json:"json" mapstructure:"json"`
}
type MessageElementJson struct {
	Type string `json:"type" mapstructure:"type"`
	Xml  string `json:"xml"  mapstructure:"xml"`
}
type MessageElementApp struct {
	Type    string `json:"type"    mapstructure:"type"`
	Content string `json:"content" mapstructure:"content"`
}

/*
name string
"Poke": 戳一戳
"ShowLove": 比心
"Like": 点赞
"Heartbroken": 心碎
"SixSixSix": 666
"FangDaZhao": 放大招
*/
type MessageElementPoke struct {
	Type string `json:"type" mapstructure:"type"`
	Name string `json:"name" mapstructure:"name"`
}
type MessageElementDice struct {
	Type  string `json:"type"  mapstructure:"type"`
	Value int    `json:"value" mapstructure:"value"`
}
type MessageElementMarkectFace struct {
	Type string `json:"type" mapstructure:"type"`
	Id   int    `json:"id"   mapstructure:"id"`
	Name string `json:"name" mapstructure:"name"`
}
type MessageElementMusicShare struct {
	Type       string `json:"type"       mapstructure:"type"`
	Kind       string `json:"kind"       mapstructure:"kind"`
	Title      string `json:"title"      mapstructure:"title"`
	Summary    string `json:"summary"    mapstructure:"summary"`
	JumpUrl    string `json:"jumpUrl"    mapstructure:"jumpUrl"`
	PictureUrl string `json:"pictureUrl" mapstructure:"pictureUrl"`
	MusicUrl   string `json:"musicUrl"   mapstructure:"musicUrl"`
	Brief      string `json:"brief"      mapstructure:"brief"`
}
type MessageElementForwardMessage struct {
	Type     string `json:"type" mapstructure:"type"`
	NodeList []struct {
		SenderId     int64                    `json:"type"         mapstructure:"type"`
		Time         int64                    `json:"time"         mapstructure:"time"`
		SenderName   string                   `json:"senderName"   mapstructure:"senderName"`
		MessageChain []map[string]interface{} `json:"messageChain" mapstructure:"messageChain"`
		MessageId    int64                    `json:"messageId"    mapstructure:"messageId"`
		MessageRef   struct {
			MessageId int   `json:"messageId" mapstructure:"messageId"`
			Target    int64 `json:"target"    mapstructure:"target"`
		} `json:"messageRef" mapstructure:"messageRef"`
	} `json:"nodeList" mapstructure:"nodeList"`
}
type MessageElementFile struct {
	Type string `json:"type" mapstructure:"type"`
	Id   string `json:"id"   mapstructure:"id"`
	Name string `json:"name" mapstructure:"name"`
	Size uint64 `json:"size" mapstructure:"size"`
}
