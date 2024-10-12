package logic

import (
	"context"

	"github.com/MrLeeang/my-zero/loginsvc/internal/svc"
	"github.com/MrLeeang/my-zero/loginsvc/loginsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *loginsvc.LoginReq) (*loginsvc.LoginResp, error) {
	// todo: add your logic here and delete this line

	l.Logger.Info(in.Username)
	l.Logger.Info(in.Password)

	l.Logger.Info("登录接口执行中...")

	return &loginsvc.LoginResp{}, nil
}
