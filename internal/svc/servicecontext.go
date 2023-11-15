package svc

import (
	"context"

	"github.com/giteexx/codegrpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

func (s *ServiceContext) Login(ctx context.Context, api, username, password string) (token string, err error) {
	return "", nil
}

func (s *ServiceContext) TokenErrWarp(ctx context.Context, username, msg string) {

}
