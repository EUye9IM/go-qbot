package api

type Message interface {
	FriendMessage | GroupMessage
}

type FriendMessage struct {
	Type   string
	Sender struct {
		Id       int64
		NickName string
		remark   string
	}
	MessageChain []map[string]interface{}
}

type GroupMessage struct {
	Type   string
	Sender struct {
		Id                int64
		MemberName        string
		SpecialTitle      string
		Permission        string
		JoinTimestamp     int64
		MuteTimeRemaining int64
		Group             struct {
			Id         int64
			Name       string
			Permission string
		}
	}
	MessageChain []map[string]interface{}
}
