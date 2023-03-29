package base

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-example-mongo/internal/svc"
	"github.com/suyuan32/simple-admin-example-mongo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	err = l.svcCtx.Mongo.Database("simple_admin").CreateCollection(context.Background(), "school")
	if err != nil {
		return nil, errorx.NewCodeInternalError(err.Error())
	}

	return &types.BaseMsgResp{
		Msg: i18n.Success,
	}, nil
}
