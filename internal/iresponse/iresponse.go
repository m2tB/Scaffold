package iresponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应结构体
type Response struct {
	Code  Code        `json:"code"`  // 自定义错误码
	Mess  interface{} `json:"mess"`  // 信息
	Data  interface{} `json:"data"`  // 数据
	Trace string      `json:"trace"` // 链路请求ID
}

// SuccessResp 成功响应 - 自定义全部内容
func SuccessResp(ctx *gin.Context, mess interface{}, data interface{}) {
	trace, _ := ctx.Get("trace")
	ctx.JSON(http.StatusOK, &Response{
		Code:  Success,
		Mess:  mess,
		Data:  data,
		Trace: trace.(string),
	})
}

// SuccessRespForFile 成功响应 - 返回不同类型文件
func SuccessRespForFile(ctx *gin.Context, contentType string, data []byte) {
	ctx.Data(http.StatusOK, contentType, data)
}

// ErrorResp 失败响应 - 自定义全部内容
func ErrorResp(ctx *gin.Context, code Code, mess interface{}, data interface{}) {
	trace, _ := ctx.Get("trace")
	ctx.JSON(http.StatusOK, &Response{
		Code:  code,
		Mess:  mess,
		Data:  data,
		Trace: trace.(string),
	})
}
