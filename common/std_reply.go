package common

import (
	"encoding/json"
)

type StdReply struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
	ErrMsg string      `json:"err_msg"`
}

var ResultSuccess = "success"
var ResultFail = "fail"

func ChangeData2String(data interface{}) string {
	if data == nil {
		return ""
	}
	d, _ := json.Marshal(data)
	return string(d)
}
