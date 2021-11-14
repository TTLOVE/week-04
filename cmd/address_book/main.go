package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"week-04/pkg/myGrpc"

	"github.com/kataras/iris/v12"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func newApp(gs *myGrpc.Server) *iris.Application {
	app := iris.New()

	listen, err := net.Listen("tcp", "127.0.0.1:9527")
	if err != nil {
		log.Fatalf("tcp listen failed:%v", err)
	}

	// gs.Listen(listen)
	go gs.Server.Serve(listen)

	return app
}

func main() {
	// c := config.New(
	// 	config.WithSource(
	// 		file.NewSource(flagconf),
	// 	),
	// )
	// var bc conf.Bootstrap
	// if err := c.Scan(&bc); err != nil {
	// 	panic(err)
	// }

	app, cleanup, err := initApp()
	if err != nil {
		fmt.Println("启动失败", err)
	}
	defer cleanup()

	fmt.Println(app)
	// 处理rpc接口
	// go gs.Server.

	app.Run(iris.Addr(fmt.Sprintf(":%d", 8080)), iris.WithoutServerError(iris.ErrServerClosed))
}
