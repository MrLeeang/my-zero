// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: loginsvc.proto

package loginsvcclient

import (
	"context"

	"github.com/MrLeeang/my-zero/loginsvc/loginsvc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = loginsvc.Request
	Response = loginsvc.Response

	LoginReq = loginsvc.LoginReq
	LoginResp = loginsvc.LoginResp

	Loginsvc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	}

	defaultLoginsvc struct {
		cli zrpc.Client
	}
)

func NewLoginsvc(cli zrpc.Client) Loginsvc {
	return &defaultLoginsvc{
		cli: cli,
	}
}

func (m *defaultLoginsvc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := loginsvc.NewLoginsvcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultLoginsvc) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := loginsvc.NewLoginsvcClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}