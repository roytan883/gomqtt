package global

//easyjson:json
type Messages struct {
	Compress int    `json:"compress"`
	Data     []byte `json:"data"`
}

//easyjson:json
type JsonMsg struct {
	FAcc   string `json:"facc"`
	FTopic string `json:"ftopic"`
	Type   int    `json:"type"`
	Qos    int    `json:"qos"`
	Time   int    `json:"time"`
	Nick   string `json:"nick"`
	MsgID  string `json:"msgid"`
	Msg    []byte `json:"msg"`
}

//easyjson:json
type C2SMsg struct {
	ToAcc string `json:"toacc"`
	Type  int    `json:"type"`
	Qos   int    `json:"qos"`
	MsgID string `json:"msgid"`
	Msg   []byte `json:"msg"`
}