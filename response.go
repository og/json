package ogjson

type ResponseAuthCode struct {
	Code string
	Msg string
}
type ResponseFailCode struct {
	Code string
	Msg string
}
type Response struct {
	Type string `json:"type"`
	Data interface{} `json:"data"`
	Code string `json:"code"`
	Msg string `json:"msg"`
}

func EmptyObject () map[string]interface{} {
	return map[string]interface{}{}
}
func Pass(data interface{}) Response {
	return Response{
		Type: "pass",
		Data: data,
		Code: "",
		Msg: "",
	}
}
func Auth(code ResponseAuthCode) Response {
	return Response{
		Type: "auth",
		Data: EmptyObject(),
		Code: code.Code,
		Msg: code.Msg,
	}
}

func FailMsg(message string) Response {
	return Response{
		Type: "fail",
		Data: EmptyObject(),
		Code: "",
		Msg: message,
	}
}
func FailCode(code ResponseFailCode) Response {
	return Response{
		Type: "fail",
		Data: EmptyObject(),
		Code: code.Code,
		Msg: code.Msg,
	}
}

func Fail(data interface{}, code ResponseFailCode) Response {
	return Response{
		Type: "fail",
		Data: data,
		Code: code.Code,
		Msg: code.Msg,
	}
}