package logic

import (
	"context"

	"github.com/MrLeeang/my-zero/rpc/internal/svc"
	"github.com/MrLeeang/my-zero/rpc/pb"

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

func (l *PingLogic) Ping(in *pb.Placeholder) (*pb.Placeholder, error) {
	// todo: add your logic here and delete this line

	return &pb.Placeholder{}, nil
}
