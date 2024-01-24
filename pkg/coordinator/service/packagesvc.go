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
	layout := "02-01-2006"

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
	pkg.Images = p.Image
	pkg.MaxCapacity = int(p.MaxCapacity)
	pkg.Name = p.Packagename
	pkg.Availablespace = int(p.MaxCapacity)
	pkg.NumOfDestination = int(p.DestinationCount)
	pkg.MinPrice = int(p.Price)
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
		Status:  "success",
		Message: "package creation done",
	}, nil
}

func (c *CoordinatorSVC) AvailablePackageSvc(p *cpb.View) (*cpb.PackagesResponce, error) {
	var pkgs []*cpb.Package
	if p.Status == "" {
		offset := 10 * (p.Page - 1)
		limit := 10
		packages, err := c.Repo.FetchAllPackages(int(offset), limit)
		if err != nil {
			return &cpb.PackagesResponce{
				Packages: nil,
			}, errors.New("error while fetching all packages")
		}

		for _, pkges := range *packages {
			var pkg cpb.Package

			pkg.PackageId = int64(pkges.ID)
			pkg.Destination = pkges.Destination
			pkg.DestinationCount = int64(pkges.NumOfDestination)
			pkg.Enddate = pkges.EndDate.Format("02-01-2006")
			pkg.Image = pkges.Images
			pkg.Packagename = pkges.Name
			pkg.AvailableSpace = int64(pkges.Availablespace)
			pkg.Price = int64(pkges.MinPrice)
			pkg.Startdate = pkges.EndDate.Format("02-01-2006")
			pkg.Startlocation = pkges.StartLocation
			pkg.Description = pkges.Description
			pkg.MaxCapacity = int64(pkges.MaxCapacity)

			pkgs = append(pkgs, &pkg)
		}
	} else {
		offset := 10 * (p.Page - 1)
		limit := 10
		packages, err := c.Repo.FetchPackages(int(offset), limit, p.Status)
		if err != nil {
			return &cpb.PackagesResponce{
				Packages: nil,
			}, errors.New("error while fetching active packages")
		}

		for _, pkges := range packages {
			var pkg cpb.Package

			pkg.PackageId = int64(pkges.ID)
			pkg.Destination = pkges.Destination
			pkg.DestinationCount = int64(pkges.NumOfDestination)
			pkg.Enddate = pkges.EndDate.Format("02-01-2006")
			pkg.Image = pkges.Images
			pkg.Packagename = pkges.Name
			pkg.AvailableSpace = int64(pkges.Availablespace)
			pkg.Price = int64(pkges.MinPrice)
			pkg.Startdate = pkges.EndDate.Format("02-01-2006")
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
		return &cpb.Package{}, errors.New("error while fetching package")
	}

	ctgry := &cpb.Category{
		CategoryName: pkg.Category.Category,
	}

	destinations, err := c.Repo.FetchPackageDestination(pkg.ID)
	if err != nil {
		return &cpb.Package{}, errors.New("error while fetching package destination")
	}

	var dstn = []*cpb.Destination{}
	for _, dsn := range destinations {
		var ds = cpb.Destination{}
		ds.Description = dsn.Description
		ds.DestinationName = dsn.DestinationName
		ds.Image = dsn.Image
		ds.DestinationId = int64(dsn.Model.ID)
		dstn = append(dstn, &ds)
	}

	return &cpb.Package{
		Packagename:      pkg.Name,
		Startlocation:    pkg.StartLocation,
		Startdate:        pkg.StartDate.Format("02-01-2006"),
		Enddate:          pkg.EndDate.Format("02-01-2006"),
		Price:            int64(pkg.MinPrice),
		Image:            pkg.Images,
		DestinationCount: int64(pkg.NumOfDestination),
		Destination:      pkg.Destination,
		PackageId:        int64(pkg.ID),
		AvailableSpace:   int64(pkg.Availablespace),
		Description:      pkg.Description,
		MaxCapacity:      int64(pkg.MaxCapacity),
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
		}, errors.New("error while updating package status")
	}
	return &cpb.Responce{
		Status:  "success",
		Message: "package status updated",
	}, nil
}
