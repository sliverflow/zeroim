package svc

import (
	"github.com/SliverFlow/zeroim/server/app/imapi/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
