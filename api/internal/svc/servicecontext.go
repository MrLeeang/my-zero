package svc

import (
	"github.com/MrLeeang/my-zero/api/internal/config"
	"github.com/MrLeeang/my-zero/loginsvc/loginsvcclient"
	"github.com/MrLeeang/my-zero/usersvc/usersvcclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	LoginSvc loginsvcclient.Loginsvc
	UserSvc  usersvcclient.Usersvc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		LoginSvc: loginsvcclient.NewLoginsvc(zrpc.MustNewClient(c.LoginSvc)),
		UserSvc:  usersvcclient.NewUsersvc(zrpc.MustNewClient(c.UserSvc)),
	}
}
