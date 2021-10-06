package initialize

import (
	"GinNaiveAdmin/global"
	"fmt"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	global.GNA_LOG.Info("初始化Redis")
	redisCfg := global.GNA_CONF.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GNA_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
		panic(fmt.Sprintf("redis connect ping failed, err: %s", err.Error()))
	} else {
		global.GNA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		return client
	}
}
