package base

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-example-rpc/internal/model"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *example.Empty) (*example.BaseResp, error) {
	data := &model.Api{
		Model:       gorm.Model{},
		Path:        "/api/create",
		Description: "hello",
		ApiGroup:    "/apis",
		Method:      "POST",
	}
	result := l.svcCtx.DB.Create(&data)
	fmt.Println(result)

	return &example.BaseResp{}, nil
}
