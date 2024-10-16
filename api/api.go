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

var prefix = "/api"

var configFile = flag.String("f", "etc/api.yaml", "the config file")
var cert = flag.String("cert", "cert", "the cert dir")

func StartServer() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c, cert)

	// 登录接口注册
	login.RegisterHandlers(server, ctx, prefix)

	// 路由注册
	var routers = []rest.Route{}

	routers = append(routers, user.RegisterHandlers(server, ctx, prefix)...)

	// 带权限的接口
	server.AddRoutes(
		routers,
		rest.WithPrefix(prefix),
		rest.WithJwt(ctx.Config.JwtAuth.AccessSecret),
	)
	// 打印出所有路由
	server.PrintRoutes()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func main() {
	StartServer()
}
