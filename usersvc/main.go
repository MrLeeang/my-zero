package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/MrLeeang/my-zero/usersvc/internal/config"
	"github.com/MrLeeang/my-zero/usersvc/internal/server"
	"github.com/MrLeeang/my-zero/usersvc/internal/svc"
	"github.com/MrLeeang/my-zero/usersvc/usersvc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "usersvc/etc/usersvc.yaml", "the config file")
var cert = flag.String("cert", "cert", "the cert dir")

func StartServer() {
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

	tlsCfg, err := loadTLSCredentials()
	if err != nil {
		panic(err)
	}

	s.AddOptions(grpc.Creds(tlsCfg))

	fmt.Printf("Starting usersvc server at %s...\n", c.ListenOn)
	s.Start()
}

func main() {
	StartServer()
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(filepath.Join(*cert, "server-cert.pem"), filepath.Join(*cert, "server-key.pem"))
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	c := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(c), nil
}
