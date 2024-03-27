package common

/*
设置响应,带有用户好友信息
*/
type H struct {
	Code  int         `json:"Code"`
	Msg   string      `json:"Msg,omitempty"`
	Data  interface{} `json:"Data,omitempty"`
	Rows  interface{} `json:"Rows,omitempty"`
	Total interface{} `json:"Total,omitempty"`
}
