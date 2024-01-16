package service

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorSVC) AddDestinationSVC(p *cpb.Destination) (*cpb.Responce, error) {
	var destination dom.Destination

	destination.Description = p.Description
	destination.DestinationName = p.DestinationName
	destination.Image = p.Image
	destination.MaxCapacity = int(p.MaxCapacity)
	destination.MinPrice = int(p.Minprice)
	destination.PackageID = uint(p.PackageID)

	err := c.Repo.CreateDestination(&destination)
	if err != nil {
		return &cpb.Responce{
			Status:  "failure",
			Message: "destination creation error",
		}, err
	}
	return &cpb.Responce{
		Status:  "Success",
		Message: "destination creation done",
	}, nil
}

func (c *CoordinatorSVC) ViewDestinationSvc(p *cpb.View) (*cpb.Destination, error) {
	dstn, err := c.Repo.FecthDestination(uint(p.Id))
	if err != nil {
		return &cpb.Destination{}, err
	}

	activity, err := c.Repo.FecthDestinationActivity(dstn.ID)
	if err != nil {
		return &cpb.Destination{}, err
	}

	actvt := cpb.Activity{}
	var arr []*cpb.Activity
	for _, act := range activity {
		actvt.ActivityType = act.ActivityType
		actvt.Activityname = act.ActivityName
		actvt.Amount = int64(act.Amount)
		actvt.Date = act.Date.Format("2006-01-02")
		actvt.Description = act.Description
		actvt.Location = act.Location
		actvt.Time = act.Time.Format("03:04 PM")
		actvt.ActivityId = int64(act.Model.ID)

		arr = append(arr, &actvt)
	}

	return &cpb.Destination{
		DestinationId:   int64(dstn.ID),
		DestinationName: dstn.DestinationName,
		Description:     dstn.Description,
		Minprice:        int64(dstn.MinPrice),
		MaxCapacity:     int64(dstn.MaxCapacity),
		Image:           dstn.Image,
		Activity:        arr,
	}, nil
}