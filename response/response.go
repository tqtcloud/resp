package response

import (
	"net/http"

	"github.com/tqtcloud/resp/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = err.(*errorx.CodeError).Code
		body.Msg = err.(*errorx.CodeError).Msg
	} else {
		body.Msg = "Success"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
