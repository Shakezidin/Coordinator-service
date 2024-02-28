package service

import (
	"errors"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

// AddDestinationSVC adds a new destination.
func (c *CoordinatorSVC) AddDestinationSVC(p *cpb.Destination) (*cpb.Response, error) {
	// Check if the package exists
	pkg, err := c.Repo.FetchPackage(uint(p.Package_ID))
	if err != nil {
		return &cpb.Response{
			Status:  "fail",
			Message: "package not found",
		}, errors.New("package not found")
	}

	// Check if the number of destinations exceeded
	dstn, _ := c.Repo.FetchPackageDestination(uint(p.Package_ID))
	if len(dstn) >= pkg.NumOfDestination {
		return &cpb.Response{
			Status:  "fail",
			Message: "number of destinations exceeded",
		}, errors.New("exceeded the destination limit")
	}

	// Create a new destination
	var destination dom.Destination
	destination.Description = p.Description
	destination.DestinationName = p.Destination_Name
	destination.Image = p.Image
	destination.PackageID = uint(p.Package_ID)
	destination.TransportationMode = p.Transportation_Mode
	destination.ArrivalLocation = p.Arrival_Location

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
		ID:      int64(destination.ID),
	}, nil
}

// ViewDestinationSvc retrieves information about a destination.
func (c *CoordinatorSVC) ViewDestinationSvc(p *cpb.View) (*cpb.Destination, error) {
	// Fetch the destination
	dstn, err := c.Repo.FecthDestination(uint(p.ID))
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
			Activity_Type: act.ActivityType,
			Activity_Name: act.ActivityName,
			Amount:        int64(act.Amount),
			Date:          act.Date.Format("02-01-2006"),
			Description:   act.Description,
			Location:      act.Location,
			Time:          act.Time.Format("03:04 PM"),
			Activity_ID:   int64(act.Model.ID),
		}
		arr = append(arr, &actvt)
	}

	// Prepare destination response
	return &cpb.Destination{
		Destination_ID:      int64(dstn.ID),
		Destination_Name:    dstn.DestinationName,
		Description:         dstn.Description,
		Image:               dstn.Image,
		Transportation_Mode: dstn.TransportationMode,
		Arrival_Location:    dstn.ArrivalLocation,
		Activity:            arr,
	}, nil
}
