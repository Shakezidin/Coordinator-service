package service

import (
	"errors"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

// AddActivitySVC adds a new activity.
func (c *CoordinatorSVC) AddActivitySVC(p *cpb.Activity) (*cpb.Response, error) {
	var activity dom.Activity
	layout := "02-01-2006"

	// Parse date and time
	date, err := time.Parse(layout, p.Date)
	if err != nil {
		return &cpb.Response{
			Status:  "failure",
			Message: "error while parsing date",
		}, errors.New("error while parsing date")
	}

	time, err := time.Parse("03:04 PM", p.Time)
	if err != nil {
		return &cpb.Response{
			Status:  "failure",
			Message: "error while parsing time",
		}, errors.New("error while parsing time")
	}

	// Set activity details
	activity.ActivityName = p.Activityname
	activity.ActivityType = p.ActivityType
	activity.Amount = int(p.Amount)
	activity.Date = date
	activity.Time = time
	activity.Description = p.Description
	activity.DestinationID = uint(p.DestinationId)
	activity.Location = p.Location

	// Create activity in the repository
	err = c.Repo.CreateActivity(&activity)
	if err != nil {
		return &cpb.Response{
			Status:  "failure",
			Message: "error while creating activity",
		}, errors.New("error while creating activity")
	}

	// Return success response
	return &cpb.Response{
		Status: "success",
		Id:     int64(activity.ID),
	}, nil
}

// ViewActivitySvc fetches details of an activity by ID.
func (c *CoordinatorSVC) ViewActivitySvc(p *cpb.View) (*cpb.Activity, error) {
	// Fetch activity from the repository
	activity, err := c.Repo.FetchActivity(uint(p.Id))
	if err != nil {
		return &cpb.Activity{}, errors.New("error while fetching activity")
	}

	// Return activity details
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
