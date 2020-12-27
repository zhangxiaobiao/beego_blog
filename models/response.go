package models

type R interface {
	Error()
	Success()
}

type Response struct {
	Code int
	Msg string
	Data interface{}
}

func (this *Response) Error(msg string,data interface{}) Response {
	return Response{
		Code: 1,
		Msg:  msg,
		Data: data,
	}
}

func (this *Response) Success(msg string,data interface{}) Response {
	return Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
}
