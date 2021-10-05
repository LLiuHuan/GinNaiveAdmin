package initialize

import (
	"GinNaiveAdmin/global"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// Close 优雅关机
func Close(s *http.Server) {
	// ******优雅关机******
	go func() {
		// 开启一个goroutine启动服务
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.GNA_LOG.Fatal("listen", zap.Error(err))
			fmt.Println("listen: ", "关机")
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	global.GNA_LOG.Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := s.Shutdown(ctx); err != nil {
		global.GNA_LOG.Fatal("Server Shutdown: ", zap.Error(err))
	}

	global.GNA_LOG.Info("Server exiting")
}