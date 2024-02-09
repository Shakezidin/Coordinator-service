package config

import (
	SVCinter "github.com/Shakezidin/pkg/coordinator/service/interface"
	"github.com/robfig/cron"
)

func InitCron(SVC SVCinter.CoordinatorSVCInter) {
	c := cron.New()
	c.AddFunc("0 6 * * *", SVC.FetchNextDayTrip)
	c.AddFunc("0 6 * * *", SVC.UpdateExpiredPackage)
	c.Start()
}
