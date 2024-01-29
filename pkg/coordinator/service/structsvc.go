package service

import (
	"github.com/Shakezidin/config"
	admin "github.com/Shakezidin/pkg/coordinator/client/pb"
	inter "github.com/Shakezidin/pkg/coordinator/repository/interface"
	SVCinter "github.com/Shakezidin/pkg/coordinator/service/interface"
	"github.com/go-redis/redis/v8"
)

type CoordinatorSVC struct {
	Repo   inter.CoordinatorRepoInter
	twilio *config.TwilioVerify
	redis  *redis.Client
	cfg    *config.Config
	client admin.AdminClient
}

func NewCoordinatorSVC(repo inter.CoordinatorRepoInter, twilio *config.TwilioVerify, redis *redis.Client, cfg *config.Config, client admin.AdminClient) SVCinter.CoordinatorSVCInter {
	return &CoordinatorSVC{
		Repo:   repo,
		twilio: twilio,
		redis:  redis,
		cfg:    cfg,
		client: client,
	}
}
