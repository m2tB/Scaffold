package imiddleware

import (
	"GhortLinks/internal/initialize/icommon"
	"GhortLinks/utils/izap/color"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// ZapLogger 接收gin框架默认的日志
func ZapLogger(mode string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		// logger记录的基础数据
		zapField := []zap.Field{
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
		}
		// 链路追踪 traceId
		trace := icommon.CURRENT_SNOW_GENERAL.Generate().Int64()
		ctx.Set(icommon.CURRENT_TRACE_ID, trace)
		// 请求日志记录
		zap.L().Debug(
			color.PrintColorStringForZap(zap.DebugLevel, fmt.Sprintf("链路追踪%s: %v", icommon.CURRENT_TRACE_ID, trace))+"  |",
			zapField...,
		)
		ctx.Next()

		latency := time.Now().Sub(start)
		if latency > time.Minute {
			latency = latency - latency%time.Second
		}
		if mode == "product" {
			zapField = append(zapField, zap.String("user-agent", ctx.Request.UserAgent()))
		}
		if query != "" {
			zapField = append(zapField, zap.String("query", query))
		}
		err := ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		if err != "" {
			zapField = append(zapField, zap.String("errors", err))
		}
		zapField = append(zapField, zap.Reflect("response", ctx.Writer))
		// 请求日志结果记录
		zap.L().Debug(color.PrintColorContentForZapWithTime(zap.DebugLevel, trace, ctx.Writer.Status(), ctx.ClientIP(), latency), zapField...)
	}
}
