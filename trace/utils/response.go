package utils

import (
	"encoding/json"
	"fmt"
)

type returnDataStruct struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"` //`json:map`
}

func Send_json(ret int, msg string, data interface{}) []byte {

	retu, err := json.Marshal(returnDataStruct{
		Ret:  ret,
		Msg:  msg,
		Data: data,
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	return retu
}
func Send_custom_json(ret int, data_key string, data interface{}) []byte {
	response := make(map[string]interface{})
	response["ret"] = ret
	response[data_key] = data
	retu, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error:", err)
	}
	return retu
}
