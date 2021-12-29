package web

import "strconv"

type Response struct {
	Code  string
	Data  interface{}
	Error string
}

func NewResponse(code int, data interface{}, err string) Response {
	if code < 300 {
		return Response{
			Code:  strconv.FormatInt(int64(code), 10),
			Data:  data,
			Error: "",
		}
	}

	return Response{
		Code:  strconv.FormatInt(int64(code), 10),
		Data:  nil,
		Error: err,
	}
}
