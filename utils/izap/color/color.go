package color

import (
	"GhortLinks/internal/initialize/icommon"
	"fmt"
	"go.uber.org/zap/zapcore"
	"time"
)

// countColorForZap 根据 izap level 计算当前输出日志类型的颜色值
func countColorForZap(level zapcore.Level) uint8 {
	if level > 1 {
		return 31
	}
	return uint8(level + 34)
}

// PrintColorLevelForZap 输出 izap 日志类型颜色内容
func PrintColorLevelForZap(level zapcore.Level) string {
	return fmt.Sprintf("\x1b[%dm%-7s\x1b[0m", countColorForZap(level), fmt.Sprintf("[%s]", level.CapitalString()))
}

// PrintColorStringForZap 输出 izap 自定义日志内容颜色内容
func PrintColorStringForZap(level zapcore.Level, content string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", countColorForZap(level), fmt.Sprintf("%s", content))
}

// PrintColorIntForZap 输出 izap 自定义日志数字类型颜色内容
func PrintColorIntForZap(level zapcore.Level, status int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", countColorForZap(level), fmt.Sprintf("%d", status))
}

// PrintColorContentForZapWithTime 输出 izap 自定义日志类型颜色内容 - 带响应时间
func PrintColorContentForZapWithTime(level zapcore.Level, ltUuid int64, status int, content string, latency time.Duration) string {
	return PrintColorStringForZap(zapcore.ErrorLevel, fmt.Sprintf("链路追踪%s: %v", icommon.CURRENT_TRACE_ID, ltUuid)) + fmt.Sprintf("  |  %-15s  |  ", content) + PrintColorIntForZap(level, status) + "  |  " + PrintColorStringForZap(zapcore.ErrorLevel, fmt.Sprintf("callback: %-10v", latency)) + "  |"
}
