package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type ResponseData struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Status    string      `json:"status"`
	TypeError string      `json:"typeError"`
}

func SuccessResponse(r *ghttp.Request, message string, code int, data interface{}) {
	var rsp = ResponseData{
		Message:   message,
		Data:      data,
		Status:    "Success",
		TypeError: "",
	}

	r.Response.WriteStatus(httpResponse[code].Status)
	r.Response.WriteJson(rsp)
}

func ErrorResponse(r *ghttp.Request, message string, code int) {
	var rsp = ResponseData{
		Message:   message,
		Data:      nil,
		Status:    "Error",
		TypeError: httpResponse[code].Type,
	}

	r.Response.WriteStatus(httpResponse[code].Status)
	r.Response.WriteJson(rsp)
}
