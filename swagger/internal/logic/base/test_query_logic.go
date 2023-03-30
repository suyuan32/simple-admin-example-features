package base

import (
	"context"
	"fmt"

	"swagger/internal/svc"
	"swagger/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestQueryLogic {
	return &TestQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestQueryLogic) TestQuery(req *types.QueryType) (resp *types.BaseMsgResp, err error) {
	return &types.BaseMsgResp{Msg: fmt.Sprintf("name:%s,age:%s", req.Name, req.Age)}, nil
}
