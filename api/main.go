package main

import (
	"flag"
	"fmt"

	"github.com/MrLeeang/my-zero/api/internal/config"
	"github.com/MrLeeang/my-zero/api/internal/handler/login"
	"github.com/MrLeeang/my-zero/api/internal/handler/user"
	"github.com/MrLeeang/my-zero/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func StartServer() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	login.RegisterHandlers(server, ctx)

	// 路由注册
	var routers = []rest.Route{}

	routers = append(routers, user.RegisterHandlers(server, ctx)...)

	server.AddRoutes(
		routers,
		rest.WithPrefix("/api"),
		rest.WithJwt("abc123abc123abc123"),
	)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func main() {
	StartServer()
}
