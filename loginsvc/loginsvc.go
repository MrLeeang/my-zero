package main

import (
	"flag"
	"fmt"

	"github.com/MrLeeang/my-zero/loginsvc/internal/config"
	"github.com/MrLeeang/my-zero/loginsvc/internal/server"
	"github.com/MrLeeang/my-zero/loginsvc/internal/svc"
	"github.com/MrLeeang/my-zero/loginsvc/loginsvc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/loginsvc.yaml", "the config file")

func StartServer() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		loginsvc.RegisterLoginsvcServer(grpcServer, server.NewLoginsvcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func main() {
	StartServer()
}
