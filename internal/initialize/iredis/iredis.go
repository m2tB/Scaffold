package iredis

import (
	"GhortLinks/internal/initialize/istruct"
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisPool map[int]*redis.Pool

func InitializeRedisPool(config *istruct.IStruct) error {
	// 将0号池默认分配给gin的sessions做存储地址
	for i := 0; i < config.RedisDbUse-1; i++ {
		redisPool[i] = &redis.Pool{
			//初始最大空闲连接数量
			MaxIdle: 10,
			//连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
			MaxActive: 0,
			//连接关闭时间 10秒 （10秒不使用自动关闭）
			IdleTimeout: 10 * time.Second,
			Dial: func() (redis.Conn, error) {
				redisDB, err := redis.Dial(
					"tcp",
					config.RedisHost,
					redis.DialConnectTimeout(time.Duration(config.RedisConnTimeout)*time.Millisecond),
					redis.DialReadTimeout(time.Duration(config.RedisReadTimeout)*time.Millisecond),
					redis.DialWriteTimeout(time.Duration(config.RedisWriteTimeout)*time.Millisecond),
					redis.DialDatabase(i+1),
					redis.DialPassword(config.RedisPassword),
				)
				if err != nil {
					return nil, err
				}
				return redisDB, nil
			},
		}
	}
	return nil
}

func GetRedis(redisPosition int) redis.Conn {
	maxLen := len(redisPool) - 1
	if redisPosition > maxLen {
		return redisPool[maxLen].Get()
	}
	return redisPool[redisPosition].Get()
}

func CloseRedisPool() error {
	for _, pool := range redisPool {
		if err := pool.Close(); err != nil {
			return err
		}
	}
	return nil
}
