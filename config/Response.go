package config

import (
	"encoding/json"
	"fmt"
)

type Resp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (res Resp) writeData(data interface{}) Resp {
	return Resp{
		Code: res.Code,
		Msg: res.Msg,
		Data: data,
	}
}

func Success(data interface{}) Resp{
	response := NewResponse(200, "success")
	return response.writeData(data)
}

func Error(err string) Resp {
	response := NewResponse(500, fmt.Sprintf("error:%s", err))
	return *response
}

func (res Resp) WriteMsg(msg string) Resp {
	return Resp{
		Code: res.Code,
		Msg: msg,
		Data: res.Data,
	}
}

func (res Resp) ToString() string {
	err:=&struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg: res.Msg,
		Data: res.Data,
	}
	resJson, _ := json.Marshal(err)
	return string(resJson)
}

func NewResponse(code int, msg string) *Resp {
	return &Resp{
		Code: code,
		Msg: msg,
		Data: nil,
	}
}

func Get() {

}