package main

import (
	"GinNaiveAdmin/global"
	"GinNaiveAdmin/initialize"
)

func init() {
	global.GNA_VP = initialize.Viper() // 初始化viper
	global.GNA_LOG = initialize.Zap()  // 初始化zap日志库
}

func main() {
	initialize.RunServer()

	if global.GNA_DB != nil {
		db, _ := global.GNA_DB.DB()
		defer db.Close()
	}
}
