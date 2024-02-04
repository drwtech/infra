package cron

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

type Server struct {
	Cron *cron.Cron
	log  *log.Helper
}

func NewServer(logger log.Logger) *Server {
	c := cron.New()
	return &Server{Cron: c, log: log.NewHelper(logger)}
}

func (s *Server) Start(ctx context.Context) error {
	s.Cron.Start()
	s.log.WithContext(ctx).Info("Cron Server Start!")
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.Cron.Stop()
	s.log.WithContext(ctx).Info("Cron Server Stop!")
	return nil
}
