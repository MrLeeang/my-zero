package main

import (
	"flag"
	"fmt"

	"github.com/MrLeeang/my-zero/usersvc/internal/config"
	"github.com/MrLeeang/my-zero/usersvc/internal/server"
	"github.com/MrLeeang/my-zero/usersvc/internal/svc"
	"github.com/MrLeeang/my-zero/usersvc/usersvc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/usersvc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		usersvc.RegisterUsersvcServer(grpcServer, server.NewUsersvcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
