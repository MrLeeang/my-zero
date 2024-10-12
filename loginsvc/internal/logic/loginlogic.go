package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/MrLeeang/my-zero/loginsvc/internal/svc"
	"github.com/MrLeeang/my-zero/loginsvc/loginsvc"
	"github.com/dgrijalva/jwt-go"

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

	if in.Username != "lihongwei" || in.Password != "123123" {
		// 登录失败
		return &loginsvc.LoginResp{}, fmt.Errorf("用户名或密码错误")
	}
	userUuid := "fdklajfklas"
	// 获取token
	token, err := GenJwtToken(in.AccessExpire, in.AccessSecret, map[string]interface{}{"uid": userUuid})

	if err != nil {
		return &loginsvc.LoginResp{}, err
	}

	return &loginsvc.LoginResp{Token: token, UserUuid: userUuid}, nil
}

func GenJwtToken(accessExpire int64, accessSecret string, payloads map[string]interface{}) (string, error) {

	now := time.Now().Unix()

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(accessSecret))
}
