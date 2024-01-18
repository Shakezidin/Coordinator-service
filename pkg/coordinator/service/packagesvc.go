package service

import (
	"errors"
	"fmt"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorSVC) AddPackageSVC(p *cpb.Package) (*cpb.Responce, error) {
	var pkg dom.Package
	layout := "2006-01-02"

	startdate, err := time.Parse(layout, p.Startdate)
	enddate, err := time.Parse(layout, p.Enddate)
	if err != nil {
		fmt.Println("date passing error")
		return &cpb.Responce{
			Status:  "fail",
			Message: "error while passing date",
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
		return &cpb.Responce{
			Status:  "failure",
			Message: "error while package creating ",
		}, err
	}
	return &cpb.Responce{
		Status:  "failure",
		Message: "package creation done",
	}, nil
}

func (c *CoordinatorSVC) AvailablePackageSvc(p *cpb.View) (*cpb.PackagesResponce, error) {
	var pkgs []*cpb.Package
	if p.Status == "" {
		packages, err := c.Repo.FetchAllPackages()
		if err != nil {
			return &cpb.PackagesResponce{
				Packages: nil,
			}, err
		}

		for _, pkges := range *packages {
			var pkg cpb.Package

			pkg.PackageId = int64(pkges.ID)
			pkg.Destination = pkges.Destination
			pkg.DestinationCount = int64(pkges.NumOfDestination)
			pkg.Enddate = pkges.EndDate.Format("2006-01-02")
			pkg.Endlocation = pkges.EndLoaction
			pkg.Image = pkges.Images
			pkg.Packagename = pkges.Name
			pkg.Price = int64(pkges.Price)
			pkg.Startdate = pkges.EndDate.Format("2006-01-02")
			pkg.Startlocation = pkges.StartLocation
			pkg.Description = pkges.Description
			pkg.MaxCapacity = int64(pkges.MaxCapacity)

			pkgs = append(pkgs, &pkg)
		}
	} else {
		packages, err := c.Repo.FetchPackages(p.Status)
		if err != nil {
			return &cpb.PackagesResponce{
				Packages: nil,
			}, err
		}

		for _, pkges := range *packages {
			var pkg cpb.Package

			pkg.PackageId = int64(pkges.ID)
			pkg.Destination = pkges.Destination
			pkg.DestinationCount = int64(pkges.NumOfDestination)
			pkg.Enddate = pkges.EndDate.Format("2006-01-02")
			pkg.Endlocation = pkges.EndLoaction
			pkg.Image = pkges.Images
			pkg.Packagename = pkges.Name
			pkg.Price = int64(pkges.Price)
			pkg.Startdate = pkges.EndDate.Format("2006-01-02")
			pkg.Startlocation = pkges.StartLocation
			pkg.Description = pkges.Description
			pkg.MaxCapacity = int64(pkges.MaxCapacity)

			pkgs = append(pkgs, &pkg)
		}
	}
	return &cpb.PackagesResponce{
		Packages: pkgs,
	}, nil
}

func (c *CoordinatorSVC) ViewPackageSVC(p *cpb.View) (*cpb.Package, error) {
	pkg, err := c.Repo.FetchPackage(uint(p.Id))
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

	var dstn = []*cpb.Destination{}
	for _, dsn := range destinations {
		var ds = cpb.Destination{}
		ds.Description = dsn.Description
		ds.DestinationName = dsn.DestinationName
		ds.Image = dsn.Image
		ds.MaxCapacity = int64(dsn.MaxCapacity)
		ds.DestinationId = int64(dsn.Model.ID)
		ds.Minprice = int64(dsn.MinPrice)

		dstn = append(dstn, &ds)
	}

	return &cpb.Package{
		Packagename:      pkg.Name,
		Startlocation:    pkg.StartLocation,
		Endlocation:      pkg.EndLoaction,
		Startdate:        pkg.StartDate.Format("2006-01-02"),
		Enddate:          pkg.EndDate.Format("2006-01-02"),
		Price:            int64(pkg.Price),
		Image:            pkg.Images,
		DestinationCount: int64(pkg.NumOfDestination),
		Destination:      pkg.Destination,
		PackageId:        int64(pkg.ID),
		Description:      pkg.Description,
		Category:         ctgry,
		Destinations:     dstn,
	}, nil
}

func (c *CoordinatorSVC) AdminPackageStatusSvc(p *cpb.View) (*cpb.Responce, error) {
	err := c.Repo.PackageStatusUpdate(uint(p.Id))
	if err != nil {
		return &cpb.Responce{
			Status:  "fail",
			Message: "error while updating package status",
		}, err
	}
	return &cpb.Responce{
		Status:  "success",
		Message: "package status updated",
	}, nil
}
