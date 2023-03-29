package config

import (
	"github.com/suyuan32/simple-admin-common/orm/gorm"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf gorm.Conf
}
