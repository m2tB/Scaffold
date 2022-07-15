package iroutes

import (
	"GhortLinks/internal/initialize/icommon"
	"GhortLinks/internal/initialize/istruct"
	"GhortLinks/internal/iresponse"
	"GhortLinks/utils/istring"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func ConsoleRoute(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	// 404 Handler 处理函数
	router.NoRoute(func(ctx *gin.Context) {
		iresponse.ErrorResp(ctx, iresponse.Refuse, iresponse.Refuse.Msg(ctx), nil)
	})

	// 初始化基于redis的存储引擎
	// 参数说明：
	//    第1个参数 - redis最大的空闲连接数
	//    第2个参数 - 数通信协议tcp或者udp
	//    第3个参数 - redis地址, 格式，host:port
	//    第4个参数 - redis密码
	//    第5个参数 - session加密密钥
	secret := icommon.CURRENT_SNOW_GENERAL.Generate().String()
	icommon.CURRENT_REDIS_SESSION_SECRET = secret
	store, err := sessions.NewRedisStore(10, "tcp", istruct.Conf.RedisHost, istruct.Conf.RedisPassword, istring.String2Byte(secret))
	store.Options(sessions.Options{
		MaxAge: icommon.DEFAULT_TOKEN_OR_JWT_EXPIRES,
	})
	if err != nil {
		log.Fatalf("Session - Redis持久化失败, err: %v\n", err)
	}

	return router
}
