package service

import (
	"errors"
	"fmt"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorSVC) AddPackageSVC(p *cpb.AddPackage) (*cpb.AddPackageResponce, error) {
	var pkg dom.Package
	layout := "2006-01-02"

	startdate, err := time.Parse(layout, p.Startdate)
	enddate, err := time.Parse(layout, p.Enddate)
	if err != nil {
		fmt.Println("date passing error")
		return &cpb.AddPackageResponce{
			Status: "date error",
		}, errors.New("date passing error")
	}

	pkg.CoordinatorId = uint(p.CoorinatorId)
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

func (c *CoordinatorSVC) AddDestinationSVC(p *cpb.AddDestination) (*cpb.AddDestinationResponce, error) {
	var destination dom.Destination

	destination.Description = p.Description
	destination.DestinationName = p.DestinationName
	destination.Image = p.Image
	destination.MaxCapacity = int(p.MaxCapacity)
	destination.MinPrice = int(p.Minprice)
	destination.PackageID = uint(p.PackageId)

	err := c.Repo.CreateDestination(&destination)
	if err != nil {
		return &cpb.AddDestinationResponce{
			Status: "destination creation error",
		}, err
	}
	return &cpb.AddDestinationResponce{
		Status: "Success",
	}, nil
}

func (c *CoordinatorSVC) AddActivitySVC(p *cpb.AddActivity) (*cpb.AddActivityResponce, error) {
	var activity dom.Activity
	layout := "2006-01-02"

	date, err := time.Parse(layout, p.Date)
	time, err1 := time.Parse("03:04 PM", p.Time)
	if err != nil {
		fmt.Println("date passing error")
		return &cpb.AddActivityResponce{
			Status: "date error",
		}, errors.New("date passing error")
	}
	if err1 != nil {
		fmt.Println("date passing errorrrrrr")
		return &cpb.AddActivityResponce{
			Status: "date error",
		}, errors.New("date passing error")
	}

	activity.ActivityName = p.ActivityName
	activity.ActivityType = p.Activitytype
	activity.Amount = int(p.Amount)
	activity.Date = date
	activity.Time = time
	activity.Description = p.Description
	activity.DestinationId = uint(p.DestinationId)
	activity.Location = p.Location

	err = c.Repo.CreateActivity(&activity)
	if err != nil {
		return &cpb.AddActivityResponce{
			Status: "destination creation error",
		}, err
	}
	return &cpb.AddActivityResponce{
		Status: "Success",
	}, nil
}

func (c *CoordinatorSVC) AvailablePackageSvc() (*cpb.PackagesResponce, error) {
	packages, err := c.Repo.FetchAllPackages()
	if err != nil {
		return &cpb.PackagesResponce{
			Packages: nil,
		}, err
	}

	var pkg cpb.Package
	var pkgs []*cpb.Package

	for _, pkges := range *packages {
		pkg.PackageId = int64(pkges.ID)
		pkg.Destination = pkges.Destination
		pkg.DestinationCount = int32(pkges.NumOfDestination)
		pkg.Enddatetime = pkges.EndDate.Format("2006-01-02")
		pkg.Endlocation = pkges.EndLoaction
		pkg.Image = pkges.Images
		pkg.Name = pkges.Name
		pkg.Price = int32(pkges.Price)
		pkg.Startdatetime = pkges.EndDate.Format("2006-01-02")
		pkg.Startlocation = pkges.StartLocation

		pkgs = append(pkgs, &pkg)
	}

	return &cpb.PackagesResponce{
		Packages: pkgs,
	}, nil
}

func (c *CoordinatorSVC) ViewPackageSVC(p *cpb.CoodinatorViewPackage) (*cpb.Package, error) {
	pkg, err := c.Repo.FetchPackage(uint(p.PackageId))
	if err != nil {
		return &cpb.Package{}, err
	}

	ctgry := &cpb.Category{
		CategoryName: pkg.Category.Category,
	}

	destinations, err := c.Repo.FetchPackageDestination(pkg.ID)
	if err != nil {
		return &cpb.Package{}, err
	}

	var ds = cpb.Destinations{}
	var dstn = []*cpb.Destinations{}
	for _, dsn := range destinations {
		ds.Description = dsn.Description
		ds.DestinationName = dsn.DestinationName
		ds.Image = dsn.Image
		ds.MaxCapacity = int64(dsn.MaxCapacity)
		ds.DestinationId = int32(dsn.ID)
		ds.MinPrice = int64(dsn.MinPrice)

		dstn = append(dstn, &ds)
	}

	return &cpb.Package{
		Name:             pkg.Name,
		Startlocation:    pkg.StartLocation,
		Endlocation:      pkg.EndLoaction,
		Startdatetime:    pkg.StartDate.Format("2006-01-02"),
		Enddatetime:      pkg.EndDate.Format("2006-01-02"),
		Price:            int32(pkg.Price),
		Image:            pkg.Images,
		DestinationCount: int32(pkg.NumOfDestination),
		Destination:      pkg.Destination,
		PackageId:        int64(pkg.ID),
		Description:      pkg.Description,
		Category:         ctgry,
		Destinations:     dstn,
	}, nil
}
