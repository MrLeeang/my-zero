package login

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/MrLeeang/my-zero/api/internal/svc"
	"github.com/MrLeeang/my-zero/api/internal/types"
	"github.com/MrLeeang/my-zero/loginsvc/loginsvc"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Ping() (resp *types.Resp, err error) {
	if _, err = l.svcCtx.LoginSvc.Ping(l.ctx, new(loginsvc.Request)); err != nil {
		return
	}

	resp = new(types.Resp)
	resp.Msg = "pong"

	return
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	var loginResp *loginsvc.LoginResp

	if loginResp, err = l.svcCtx.LoginSvc.Login(l.ctx, &loginsvc.LoginReq{
		Username:     req.Username,
		Password:     req.Password,
		AccessSecret: l.svcCtx.Config.JwtAuth.AccessSecret,
		AccessExpire: l.svcCtx.Config.JwtAuth.AccessExpire,
	}); err != nil {
		return
	}

	resp = new(types.LoginResp)
	resp.Msg = types.ErrorCodeMessage[types.Ok]
	resp.Data.Token = loginResp.Token
	resp.Data.UserUuid = loginResp.UserUuid

	return
}
