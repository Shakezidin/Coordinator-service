package service

import (
	"errors"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

// AddDestinationSVC adds a new destination.
func (c *CoordinatorSVC) AddDestinationSVC(p *cpb.Destination) (*cpb.Response, error) {
	// Check if the package exists
	pkg, err := c.Repo.FetchPackage(uint(p.PackageID))
	if err != nil {
		return &cpb.Response{
			Status:  "fail",
			Message: "package not found",
		}, errors.New("package not found")
	}

	// Check if the number of destinations exceeded
	dstn, _ := c.Repo.FetchPackageDestination(uint(p.PackageID))
	if len(dstn) >= pkg.NumOfDestination {
		return &cpb.Response{
			Status:  "fail",
			Message: "number of destinations exceeded",
		}, errors.New("exceeded the destination limit")
	}

	// Create a new destination
	var destination dom.Destination
	destination.Description = p.Description
	destination.DestinationName = p.DestinationName
	destination.Image = p.Image
	destination.PackageID = uint(p.PackageID)
	destination.TransportationMode = p.TransportationMode
	destination.ArrivalLocation = p.ArrivalLocation

	err = c.Repo.CreateDestination(&destination)
	if err != nil {
		return &cpb.Response{
			Status:  "failure",
			Message: "error while creating destination",
		}, errors.New("error while creating destination")
	}

	return &cpb.Response{
		Status:  "success",
		Message: "destination creation done",
		Id:      int64(destination.ID),
	}, nil
}

// ViewDestinationSvc retrieves information about a destination.
func (c *CoordinatorSVC) ViewDestinationSvc(p *cpb.View) (*cpb.Destination, error) {
	// Fetch the destination
	dstn, err := c.Repo.FecthDestination(uint(p.Id))
	if err != nil {
		return nil, errors.New("error while fetching destination")
	}

	// Fetch activities related to the destination
	activity, err := c.Repo.FecthDestinationActivity(dstn.ID)
	if err != nil {
		return nil, errors.New("error while fetching activities")
	}

	// Prepare activity response
	var arr []*cpb.Activity
	for _, act := range activity {
		actvt := cpb.Activity{
			ActivityType: act.ActivityType,
			Activityname: act.ActivityName,
			Amount:       int64(act.Amount),
			Date:         act.Date.Format("02-01-2006"),
			Description:  act.Description,
			Location:     act.Location,
			Time:         act.Time.Format("03:04 PM"),
			ActivityId:   int64(act.Model.ID),
		}
		arr = append(arr, &actvt)
	}

	// Prepare destination response
	return &cpb.Destination{
		DestinationId:      int64(dstn.ID),
		DestinationName:    dstn.DestinationName,
		Description:        dstn.Description,
		Image:              dstn.Image,
		TransportationMode: dstn.TransportationMode,
		ArrivalLocation:    dstn.ArrivalLocation,
		Activity:           arr,
	}, nil
}
