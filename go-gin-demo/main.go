package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/mirkowu/go-gin-demo/pkg/logging"
	"github.com/mirkowu/go-gin-demo/pkg/setting"
	"github.com/mirkowu/go-gin-demo/routers"
	"syscall"
)

func main() {

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())

	server.BeforeBegin = func(add string) {
		logging.Info("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		logging.Info("Server err: %v", err)
	}

	//router := routers.InitRouter()
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()

}
