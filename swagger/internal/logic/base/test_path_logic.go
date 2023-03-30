package base

import (
	"context"
	"fmt"

	"swagger/internal/svc"
	"swagger/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestPathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestPathLogic {
	return &TestPathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestPathLogic) TestPath(req *types.PathType) (resp *types.BaseMsgResp, err error) {
	return &types.BaseMsgResp{Msg: fmt.Sprintf("name:%s,age:%s", req.Name, req.Age)}, nil
}
