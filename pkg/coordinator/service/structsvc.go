package service

import (
	"github.com/Shakezidin/config"
	inter "github.com/Shakezidin/pkg/coordinator/repository/interface"
	SVCinter "github.com/Shakezidin/pkg/coordinator/service/interface"
	"github.com/go-redis/redis/v8"
)

type CoordinatorSVC struct {
	Repo   inter.CoordinatorRepoInter
	twilio *config.TwilioVerify
	redis  *redis.Client
	cfg    *config.Config
}

func NewCoordinatorSVC(repo inter.CoordinatorRepoInter, twilio *config.TwilioVerify, redis *redis.Client, cfg *config.Config) SVCinter.CoordinatorSVCInter {
	return &CoordinatorSVC{
		Repo:   repo,
		twilio: twilio,
		redis:  redis,
		cfg:    cfg,
	}
}
