package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go_example/web_app/Dao/mysql"
	"go_example/web_app/Dao/redis"
	"go_example/web_app/logger"
	"go_example/web_app/router"
	"go_example/web_app/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Go Web开发较为通用的脚手架模板

func main() {
	//1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}

	//2. 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// 刷新
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	//3. 初始化MySQL连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	//4. 初始化Redis连接
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	//5. 注册路由
	r := router.Setup()

	//6. 启动服务（优雅关闭)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("app.port")),
		Handler: r,
	}

	// 否则陷入死循环，在处理listenAndServe
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			zap.L().Fatal("listen:%s\n", zap.Error(err))
		}
	}()
	zap.L().Info("Server exiting")

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

}
