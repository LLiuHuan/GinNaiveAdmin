package initialize

import (
	"GinNaiveAdmin/core"
	"GinNaiveAdmin/global"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

func RunServer() *http.Server {

	// 初始化Validator翻译器
	if err := InitTrans("zh"); err != nil {
		fmt.Printf("init validator failed, err: %v\n", err)
		global.GNA_LOG.Error("init validator failed, err", zap.Error(err))
		os.Exit(0)
	}

	// 注册路由
	router := Routers(global.GNA_CONF.System.Mode)

	address := fmt.Sprintf(":%d", global.GNA_CONF.System.Port)
	s := core.InitServer(address, router)

	time.Sleep(10 * time.Microsecond)

	global.GNA_LOG.Info("服务启动成功，监听端口为：  ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 GinNaiveAdmin
	当前版本:V0.0.1
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	默认后端API运行地址:http://127.0.0.1:%s/v1/
	`, address, address)
	fmt.Println()

	Close(s)

	return s
}
