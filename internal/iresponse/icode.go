package iresponse

import (
	"GhortLinks/internal/initialize/icommon"
	"github.com/gin-gonic/gin"
)

type Code int64

// [System] 系统级统一状态码
// [Database] 数据库级统一状态码
// [Parameter] 参数库级统一状态码
const (
	System  Code = 10001
	Success      = System + iota
	Pause
	Mistake
	Limit
	Refuse
)

var codeMsgMap = map[Code]map[string]string{
	Success: {
		"zh": "正常响应",
		"en": "Response Success",
	},
	Pause: {
		"zh": "暂停响应",
		"en": "Response Pause",
	},
	Mistake: {
		"zh": "错误响应",
		"en": "Response Mistake",
	},
	Limit: {
		"zh": "访问受限",
		"en": "Response Limit",
	},
	Refuse: {
		"zh": "访问拒绝",
		"en": "Response Refuse",
	},
}

func (c Code) Msg(ctx *gin.Context) string {
	return codeMsgMap[c][ctx.DefaultQuery("locale", icommon.DEFAULT_LOCAL)]
}
