package config

import (
	SVCinter "github.com/Shakezidin/pkg/coordinator/service/interface"
	"github.com/robfig/cron"
)

// InitCron initializes and starts cron jobs for scheduling tasks.
func InitCron(SVC SVCinter.CoordinatorSVCInter) {
	// Create a new cron instance
	c := cron.New()

	// Add cron jobs
	c.AddFunc("0 6 * * *", SVC.FetchNextDayTrip)

	c.AddFunc("0 6 * * *", SVC.UpdateExpiredPackage)
	// Start the cron scheduler
	c.Start()
}
