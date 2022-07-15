package initialize

import (
	"GhortLinks/internal/initialize/icommon"
	"GhortLinks/internal/initialize/idatebase"
	"GhortLinks/internal/initialize/ilogger"
	"GhortLinks/internal/initialize/iredis"
	"GhortLinks/internal/initialize/istruct"
	"GhortLinks/utils/inet"
	"GhortLinks/utils/isnow"
	"go.uber.org/zap"
	"log"
	"time"
)

func InitForProject() {
	// 启动流程1: 加载配置数据
	if err := istruct.InitializeConfig(); err != nil {
		log.Fatalf("初始化配置信息内容失败, err: %v\n", err)
		return
	}
	// 启动流程2: 初始化项目统一全局数据
	ips := inet.GetLocationIP()
	if len(ips) > 0 {
		icommon.CURRENT_LOCAL_IP = ips[0]
	}
	if location, err := time.LoadLocation(istruct.Conf.TimeLocation); err != nil {
		log.Fatalf("初始化本地时区失败, err: %v\n", err)
		return
	} else {
		icommon.CURRENT_TIME_LOCATION = location
	}
	if snow, err := isnow.InitializeSnow(istruct.Conf.SnowRuntime, istruct.Conf.MachineID); err != nil {
		log.Fatalf("初始化snow算法失败, err: %v\n", err)
		return
	} else {
		icommon.CURRENT_SNOW_GENERAL = snow
	}
	// 启动流程3: 初始化日志框架
	if err := ilogger.InitializeLogger(istruct.Conf); err != nil {
		log.Fatalf("初始化日志框架失败, err: %v\n", err)
		return
	}
	// 启动流程4: 初始化数据库及redis
	if err := idatebase.InitializeDatabase(istruct.Conf); err != nil {
		log.Fatalf("初始化数据库连接池失败, err: %v\n", err)
		return
	}
	if err := iredis.InitializeRedisPool(istruct.Conf); err != nil {
		log.Fatalf("初始化Redis连接池失败, err: %v\n", err)
		return
	}
}

func CloseForProject() {
	_ = idatebase.CloseDatabase()
	zap.L().Debug("关闭数据库连接池成功 ...")
	_ = iredis.CloseRedisPool()
	zap.L().Debug("关闭Redis连接池成功 ...")
}
