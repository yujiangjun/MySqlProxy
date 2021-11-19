package main

import "encoding/json"

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (res Response) WriteData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg: res.Msg,
		Data: data,
	}
}

func (res Response) WriteMsg(msg string) Response {
	return Response{
		Code: res.Code,
		Msg: msg,
		Data: res.Data,
	}
}

func (res Response) ToString() string {
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

func NewResponse(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg: msg,
		Data: nil,
	}
}

func Get() {

}