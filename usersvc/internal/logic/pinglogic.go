package logic

import (
	"context"

	"github.com/MrLeeang/my-zero/usersvc/internal/svc"
	"github.com/MrLeeang/my-zero/usersvc/usersvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *usersvc.Request) (*usersvc.Response, error) {
	// todo: add your logic here and delete this line

	return &usersvc.Response{}, nil
}
