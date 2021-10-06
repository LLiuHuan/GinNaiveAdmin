package main

import (
	"GinNaiveAdmin/global"
	"GinNaiveAdmin/initialize"
)

func init() {
	global.GNA_VP = initialize.Viper() // 初始化viper
	global.GNA_LOG = initialize.Zap()  // 初始化zap日志库
	global.GNA_DB = initialize.Gorm()  // gorm连接数据库
	// 初始化Redis
	if global.GNA_CONF.System.UseMultipoint || global.GNA_CONF.RateLimit.IpVerify {
		// 初始化redis服务
		global.GNA_REDIS = initialize.Redis() // 初始化redis
	}
}

func main() {
	initialize.RunServer()

	if global.GNA_DB != nil {
		db, _ := global.GNA_DB.DB()
		defer db.Close()
	}
}
