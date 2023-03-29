package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-example-rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := c.DatabaseConf.NewGORM()
	logx.Must(err)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
