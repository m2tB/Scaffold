package iserver

import (
	"GhortLinks/internal/imiddleware"
	"GhortLinks/internal/initialize/istruct"
	"GhortLinks/internal/iroutes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var HttpConsoleSrvHandler *http.Server

func HttpConsoleServerRun() {
	// 初始默认中间件
	middleware := []gin.HandlerFunc{
		imiddleware.ZapLogger(istruct.Conf.RuntimeMode),
		imiddleware.ZapRecovery(true),
		imiddleware.Translator(),
	}
	router := iroutes.ConsoleRoute(middleware...)
	HttpConsoleSrvHandler = &http.Server{
		Addr:           istruct.Conf.HttpListenPort,
		Handler:        router,
		ReadTimeout:    time.Duration(istruct.Conf.HttpReadsTimeout) * time.Second,
		WriteTimeout:   time.Duration(istruct.Conf.HttpWriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(istruct.Conf.HttpMaxHeaderBytes),
	}
	go func() {
		zap.L().Debug(fmt.Sprintf("项目启动成功, 启动端口: %s", istruct.Conf.HttpListenPort))
		if err := HttpConsoleSrvHandler.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal(fmt.Sprintf("项目启动失败, 启动端口: %s, err: %v\n", istruct.Conf.HttpListenPort, err))
		}
	}()
}

func HttpConsoleServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpConsoleSrvHandler.Shutdown(ctx); err != nil {
		zap.L().Fatal(fmt.Sprintf("项目关闭失败, err: %v\n", err))
		return
	}
	zap.L().Debug(fmt.Sprintf("项目关闭成功, 关闭端口: %s", istruct.Conf.HttpListenPort))
}
