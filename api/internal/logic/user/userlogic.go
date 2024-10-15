package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/MrLeeang/my-zero/api/internal/svc"
	"github.com/MrLeeang/my-zero/api/internal/types"
	"github.com/MrLeeang/my-zero/usersvc/usersvc"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) Ping() (resp *types.Resp, err error) {
	if _, err = l.svcCtx.UserSvc.Ping(l.ctx, new(usersvc.Request)); err != nil {
		return
	}

	resp = new(types.Resp)
	resp.Msg = "pong"
	resp.Code = types.Ok

	return
}

// 创建用户接口
func (l *UserLogic) CreateUser(req types.CreateUserReq) (resp *types.CreateUserResp, err error) {

	l.Logger.Info(l.ctx.Value("uid"))

	in := usersvc.CreateUserReq{
		Username: req.Username,
		Password: req.Password,
	}

	var out *usersvc.CreateUserResp

	if out, err = l.svcCtx.UserSvc.CreateUser(l.ctx, &in); err != nil {
		return
	}

	resp = new(types.CreateUserResp)
	resp.Msg = types.ErrorCodeMessage[types.Ok]
	resp.Data.UserUuid = out.UserUuid
	resp.Data.Username = out.Username

	return
}
