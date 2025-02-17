package logic

import (
	"context"

	"github.com/SliverFlow/zeroim/server/app/imrpc/imrpc"
	"github.com/SliverFlow/zeroim/server/app/imrpc/internal/svc"

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

func (l *PingLogic) Ping(in *imrpc.Request) (*imrpc.Response, error) {
	// todo: add your logic here and delete this line

	return &imrpc.Response{}, nil
}
