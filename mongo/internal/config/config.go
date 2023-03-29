package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"

	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/suyuan32/simple-admin-common/plugins/storage/mongodb"
)

type Config struct {
	rest.RestConf
	Auth        rest.AuthConf
	MongodbConf mongodb.Conf
	RedisConf   redis.RedisConf
	CasbinConf  casbin.CasbinConf
}
