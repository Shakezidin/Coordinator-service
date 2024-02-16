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
	activity.ActivityName = p.Activity_Name
	activity.ActivityType = p.Activity_Type
	activity.Amount = int(p.Amount)
	activity.Date = date
	activity.Time = time
	activity.Description = p.Description
	activity.DestinationID = uint(p.Destination_ID)
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
		ID:     int64(activity.ID),
	}, nil
}

// ViewActivitySvc fetches details of an activity by ID.
func (c *CoordinatorSVC) ViewActivitySvc(p *cpb.View) (*cpb.Activity, error) {
	// Fetch activity from the repository
	activity, err := c.Repo.FetchActivity(uint(p.ID))
	if err != nil {
		return &cpb.Activity{}, errors.New("error while fetching activity")
	}

	// Return activity details
	return &cpb.Activity{
		Activity_ID:   int64(activity.ID),
		Activity_Name: activity.ActivityName,
		Description:  activity.Description,
		Location:     activity.Location,
		Activity_Type: activity.ActivityType,
		Amount:       int64(activity.Amount),
		Time:         activity.Time.Format("03:04 PM"),
		Date:         activity.Date.Format("02-01-2006"),
	}, nil
}
