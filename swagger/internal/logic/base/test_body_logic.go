package base

import (
	"context"
	"fmt"

	"swagger/internal/svc"
	"swagger/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestBodyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestBodyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestBodyLogic {
	return &TestBodyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestBodyLogic) TestBody(req *types.BodyType) (resp *types.BaseMsgResp, err error) {
	return &types.BaseMsgResp{Msg: fmt.Sprintf("name:%s,age:%s", req.Name, req.Age)}, nil
}
