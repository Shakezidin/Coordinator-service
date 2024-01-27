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
			Status: "fail",
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
			Status: "failure",
		}, errors.New("error while creating package")
	}
	return &cpb.Responce{
		Status:  "success",
		Message: "package creation done",
		Id:      int64(pkg.ID),
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
			Status: "fail",
		}, errors.New("error while updating package status")
	}
	return &cpb.Responce{
		Status:  "success",
		Message: "package status updated",
	}, nil
}

func (c *CoordinatorSVC) FilterPackageSvc(p *cpb.Filter) (*cpb.PackagesResponce, error) {
	query := "SELECT * FROM packages WHERE 1 = 1"
	var args []interface{}

	if p.Departurtime != "" {
		query += " AND start_time = ?"
		args = append(args, p.Departurtime)
	}
	if p.MinPrice > 0 {
		query += " AND min_price >= ?"
		args = append(args, p.MinPrice)
	}
	if p.MaxPrice > 0 {
		query += " AND min_price <= ?"
		args = append(args, p.MaxPrice)
	}
	if p.CategoryId != 0 {
		query += " AND trip_category_id = ?"
		args = append(args, p.CategoryId)
	}

	// Add ORDER BY clause
	if p.OrderBy != "" {
		query += " ORDER BY " + p.OrderBy
	}

	// Add LIMIT and OFFSET clauses for paging
	offset := 10 * (p.Page - 1)
	limit := 10
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	db := c.Repo.GetDB()
	rows, err := db.Raw(query, args...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []*dom.Package
	for rows.Next() {
		var p dom.Package
		if err := rows.Scan(&p.ID, &p.Name, &p.StartTime, &p.MinPrice, &p.TripCategoryId); err != nil {
			return nil, err
		}
		packages = append(packages, &p)
	}

	// Convert packages to protobuf format
	var pkgs []*cpb.Package
	for _, pkge := range packages {
		pkg := &cpb.Package{
			PackageId:        int64(pkge.ID),
			Destination:      pkge.Destination,
			DestinationCount: int64(pkge.NumOfDestination),
			Enddate:          pkge.EndDate.Format("02-01-2006"),
			Image:            pkge.Images,
			Packagename:      pkge.Name,
			AvailableSpace:   int64(pkge.Availablespace),
			Price:            int64(pkge.MinPrice),
			Startdate:        pkge.StartDate.Format("02-01-2006"),
			Startlocation:    pkge.StartLocation,
			Description:      pkge.Description,
			MaxCapacity:      int64(pkge.MaxCapacity),
		}
		pkgs = append(pkgs, pkg)
	}

	return &cpb.PackagesResponce{
		Packages: pkgs,
	}, nil
}
