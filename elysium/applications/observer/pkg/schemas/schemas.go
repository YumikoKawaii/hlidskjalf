package schemas

import "encoding/json"

type Message struct {
	SessionId string `json:"sessionId,omitempty"`
	View      int64  `json:"view,omitempty"`
}

func (m *Message) ToByte() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
