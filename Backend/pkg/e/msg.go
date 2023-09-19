package e

var MsgFlag = map[int]string{
	Success: "操作成功",
	Error:   "操作失败",

	InvalidParams: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if !ok {
		return MsgFlag[Error]
	}
	return msg
}
