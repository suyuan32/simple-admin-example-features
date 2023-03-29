package base

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-example-mongo/internal/svc"
	"github.com/suyuan32/simple-admin-example-mongo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateExampleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateExampleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExampleLogic {
	return &CreateExampleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Book struct {
	Title  string
	Author string
}

func (l *CreateExampleLogic) CreateExample() (resp *types.BaseMsgResp, err error) {
	doc := Book{Title: "Atonement", Author: "Ian McEwan"}
	one, err := l.svcCtx.MongoDatabase.Collection("school").InsertOne(
		l.ctx, doc)
	if err != nil {
		return nil, errorx.NewCodeInternalError(err.Error())
	}

	fmt.Println(one)

	return &types.BaseMsgResp{Msg: i18n.Success}, nil
}
