package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/MrLeeang/my-zero/api/internal/svc"
	"github.com/MrLeeang/my-zero/api/internal/types"
	"github.com/MrLeeang/my-zero/usersvc/usersvc"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.Resp, err error) {
	if _, err = l.svcCtx.UserSvc.Ping(l.ctx, new(usersvc.Request)); err != nil {
		return
	}

	resp = new(types.Resp)
	resp.Msg = "pong"

	return
}
