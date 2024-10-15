package logic

import (
	"context"
	"fmt"

	"github.com/MrLeeang/my-zero/database"
	"github.com/MrLeeang/my-zero/loginsvc/internal/svc"
	"github.com/MrLeeang/my-zero/loginsvc/loginsvc"
	"github.com/MrLeeang/my-zero/utils"

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

	userModel := database.NewSysUserModel(database.Conn)

	user, err := userModel.FindOneByLoginUser(l.ctx, in.Username)

	if err != nil {
		return &loginsvc.LoginResp{}, fmt.Errorf("user %s is not found", in.Username)
	}

	ok := utils.ComparePassword(user.LoginPass, in.Password)

	if !ok {
		return &loginsvc.LoginResp{}, fmt.Errorf("密码错误")
	}

	// 获取token
	token, err := utils.GenJwtToken(in.AccessExpire, in.AccessSecret, map[string]interface{}{"uid": user.Uuid})

	if err != nil {
		return &loginsvc.LoginResp{}, err
	}

	return &loginsvc.LoginResp{Token: token, UserUuid: user.Uuid}, nil
}
