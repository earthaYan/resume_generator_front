package ctl

import (
	"encoding/json"
	"resume_backend/pkg/e"
)

// 最基础的返回resp
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"message"`
	Error  string      `json:"error"`
}

// 带数组形式的返回
type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

// 用户登录专属的的返回
type TokenData struct {
	User        interface{} `json:"user"`
	AccessToken string      `json:"access_token"`
}

// 带有追踪信息的错误返回
type TrackedErrorResponse struct {
	Response
	TraceId string `json:"trace_id"`
}

// 2xx 正常返回成功的函数
func RespSuccess(code ...int) *Response {
	status := e.Success
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMsg(status),
	}
}

func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := e.Success
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}
}

func RespError(err error, data interface{}, code ...int) *TrackedErrorResponse {
	status := e.Error
	if code != nil {
		status = code[0]
	}
	return &TrackedErrorResponse{
		Response: Response{
			Status: status,
			Data:   data,
			Msg:    e.GetMsg(status),
			Error:  err.Error(),
		},
		TraceId: "",
	}
}

func ErrorResponse(err error) *TrackedErrorResponse {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return RespError(err, "json类型不匹配")
	}
	return RespError(err, "参数错误", e.InvalidParams)
}

func ResponseList(items interface{}, total int64) Response {
	return Response{
		Status: e.Success,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "操作成功",
	}
}
