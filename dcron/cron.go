package dcron

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/libi/dcron"
	"github.com/libi/dcron/driver"
	"github.com/redis/go-redis/v9"
)

// https://github.com/libi/dcron/tree/master
type Server struct {
	Cron *dcron.Dcron
}

func NewServer(rds *redis.Client, serviceName string, logger log.Logger) *Server {
	drv := driver.NewRedisDriver(rds)
	dLog := &DLog{Log: log.NewHelper(logger)}
	cron := dcron.NewDcronWithOption(
		serviceName,
		drv,
		dcron.WithLogger(dLog),
		dcron.WithHashReplicas(10),
		dcron.WithNodeUpdateDuration(time.Second*10),
	)

	return &Server{Cron: cron}
}

func (s *Server) Start(ctx context.Context) error {
	s.Cron.Start()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.Cron.Stop()
	return nil
}
