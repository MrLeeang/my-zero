package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/MrLeeang/my-zero/usersvc/usersvcclient"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	config := zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"192.168.2.235:2379"},
			Key:   "usersvc.rpc",
		},
	}

	cert := "cert"

	tlsCredentials, _ := loadTLSCredentials(&cert)

	// zrpc.WithTransportCredentials(tlsCredentials)

	cli := usersvcclient.NewUsersvc(zrpc.MustNewClient(config, zrpc.WithTransportCredentials(tlsCredentials)))
	res, err := cli.Ping(context.TODO(), &usersvcclient.Request{})

	if err != nil {
		panic(err)
	}

	fmt.Println(res.Pong)
}

func loadTLSCredentials(cert *string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(*cert + "/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}
