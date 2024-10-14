package logic

import (
	"context"

	"github.com/MrLeeang/my-zero/usersvc/internal/svc"
	"github.com/MrLeeang/my-zero/usersvc/usersvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLogic) CreateUser(in *usersvc.CreateUserReq) (*usersvc.CreateUserResp, error) {
	// todo: add your logic here and delete this line

	return &usersvc.CreateUserResp{UserUuid: "xxxxx", Username: "test创建用户"}, nil
}
