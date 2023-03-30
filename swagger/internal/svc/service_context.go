package svc

import (
	"swagger/internal/config"

	"swagger/ent"
)

type ServiceContext struct {
	Config config.Config

	DB *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
	}
}
