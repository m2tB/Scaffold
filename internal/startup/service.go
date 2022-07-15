package startup

import (
	"GhortLinks/internal/initialize"
	"GhortLinks/internal/initialize/iserver"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func Service() {
	// 初始化应用配置信息
	initialize.InitForProject()
	defer initialize.CloseForProject()

	// 项目启动及关闭处理
	iserver.HttpServiceServerRun()
	// 等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Debug("正在关闭系统服务 ...")
	iserver.HttpServiceServerStop()
}
