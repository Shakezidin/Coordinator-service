package service

import (
	"errors"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorSVC) AddActivitySVC(p *cpb.Activity) (*cpb.Responce, error) {
	var activity dom.Activity
	layout := "02-01-2006"

	date, err := time.Parse(layout, p.Date)
	time, err1 := time.Parse("03:04 PM", p.Time)
	if err != nil {
		return &cpb.Responce{
			Status:  "filure",
			Message: "error while passig date",
		}, errors.New("error while passig date")
	}
	if err1 != nil {
		return &cpb.Responce{
			Status:  "failure",
			Message: "error while passing time",
		}, errors.New("error while passing time")
	}

	activity.ActivityName = p.Activityname
	activity.ActivityType = p.ActivityType
	activity.Amount = int(p.Amount)
	activity.Date = date
	activity.Time = time
	activity.Description = p.Description
	activity.DestinationId = uint(p.DestinationId)
	activity.Location = p.Location

	err = c.Repo.CreateActivity(&activity)
	if err != nil {
		return &cpb.Responce{
			Status:  "failure",
			Message: "error while destination creation",
		}, errors.New("error while destination creation")
	}
	return &cpb.Responce{
		Status:  "Success",
		Message: "destination creation done",
		Id:      int64(activity.ID),
	}, nil
}

func (c *CoordinatorSVC) ViewActivitySvc(p *cpb.View) (*cpb.Activity, error) {
	activity, err := c.Repo.FecthActivity(uint(p.Id))
	if err != nil {
		return &cpb.Activity{}, errors.New("error while fetching activity")
	}

	return &cpb.Activity{
		ActivityId:   int64(activity.ID),
		Activityname: activity.ActivityName,
		Description:  activity.Description,
		Location:     activity.Location,
		ActivityType: activity.ActivityType,
		Amount:       int64(activity.Amount),
		Time:         activity.Time.Format("03:04 PM"),
		Date:         activity.Date.Format("02-01-2006"),
	}, nil

}
