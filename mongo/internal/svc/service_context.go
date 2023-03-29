package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/suyuan32/simple-admin-example-mongo/internal/config"
	i18n2 "github.com/suyuan32/simple-admin-example-mongo/internal/i18n"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type ServiceContext struct {
	Config        config.Config
	Casbin        *casbin.Enforcer
	Authority     rest.Middleware
	Trans         *i18n.Translator
	Mongo         *mongo.Client
	MongoDatabase *mongo.Database
}

func NewServiceContext(c config.Config) *ServiceContext {

	trans := i18n.NewTranslator(i18n2.LocaleFS)

	return &ServiceContext{
		Config:        c,
		Trans:         trans,
		Mongo:         c.MongodbConf.MustNewClient(),
		MongoDatabase: c.MongodbConf.MustNewDatabase(),
	}
}
