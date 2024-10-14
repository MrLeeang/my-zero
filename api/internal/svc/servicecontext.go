package svc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/MrLeeang/my-zero/api/internal/config"
	"github.com/MrLeeang/my-zero/loginsvc/loginsvcclient"
	"github.com/MrLeeang/my-zero/usersvc/usersvcclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/credentials"
)

type ServiceContext struct {
	Config   config.Config
	LoginSvc loginsvcclient.Loginsvc
	UserSvc  usersvcclient.Usersvc
}

func NewServiceContext(c config.Config, cert *string) *ServiceContext {

	tlsCredentials, err := loadTLSCredentials(cert)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:   c,
		LoginSvc: loginsvcclient.NewLoginsvc(zrpc.MustNewClient(c.LoginSvc, zrpc.WithTransportCredentials(tlsCredentials))),
		UserSvc:  usersvcclient.NewUsersvc(zrpc.MustNewClient(c.UserSvc, zrpc.WithTransportCredentials(tlsCredentials))),
	}
}

func loadTLSCredentials(cert *string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate

	pemServerCA, err := ioutil.ReadFile(filepath.Join(*cert, "ca-cert.pem"))
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
