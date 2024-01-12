package service

import (
	"errors"
	"fmt"
	"time"

	dom "github.com/Shakezidin/pkg/DOM/coordinator"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

func (c *CoordinatorSVC) AddPackageSVC(p *cpb.AddPackage) (*cpb.AddPackageResponce, error) {
	var pkg dom.Package
	user, err := c.Repo.FindUserByEmail(p.Coorinatoremail)
	if err != nil {
		fmt.Println("coordinator fetching error")
		return &cpb.AddPackageResponce{
			Status: "user not found",
		}, err
	}

	if user.Isblock == true {
		fmt.Println("user is blocked")
		return &cpb.AddPackageResponce{
			Status: "user is blocked",
		}, errors.New("user is blocked")
	}
	layout := "2006-01-02"

	startdate, err := time.Parse(layout, p.Startdate)
	enddate, err := time.Parse(layout, p.Enddate)

	pkg.CoordinatorId = user.ID
	pkg.Description = p.Description
	pkg.Destination = p.Destination
	pkg.EndDate = enddate
	pkg.EndLoaction = p.Endlocation
	pkg.Images = p.Image
	pkg.MaxCapacity = int(p.MaxCapacity)
	pkg.Name = p.Packagename
	pkg.NumOfDestination = int(p.DestinationCount)
	pkg.Price = int(p.Price)
	pkg.StartDate = startdate
	pkg.StartLocation = p.Startlocation
	pkg.TripCategoryId = uint(p.CategoryId)

	err = c.Repo.CreatePackage(&pkg)
	if err != nil {
		return &cpb.AddPackageResponce{
			Status: "package creation error",
		}, err
	}
	return &cpb.AddPackageResponce{
		Status: "Success",
	}, nil

}
