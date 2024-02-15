package service

import (
	"errors"
	"fmt"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

// AddPackageSVC adds a new package.
func (c *CoordinatorSVC) AddPackageSVC(p *cpb.Package) (*cpb.Response, error) {
	var pkg dom.Package
	layout := "02-01-2006"

	startDate, err := time.Parse(layout, p.Startdate)
	if err != nil {
		return &cpb.Response{Status: "fail"}, errors.New("failed to parse startdate")
	}

	endDate, err := time.Parse(layout, p.Enddate)
	if err != nil {
		return &cpb.Response{Status: "fail"}, errors.New("failed to parse enddate")
	}

	pkg.CoordinatorID = uint(p.CoorinatorId)
	pkg.Description = p.Description
	pkg.Destination = p.Destination
	pkg.EndDate = endDate
	pkg.Images = p.Image
	pkg.MaxCapacity = int(p.MaxCapacity)
	pkg.Name = p.Packagename
	pkg.AvailableSpace = int(p.MaxCapacity)
	pkg.NumOfDestination = int(p.DestinationCount)
	pkg.MinPrice = int(p.Price)
	pkg.StartDate = startDate
	pkg.StartTime = p.Starttime
	pkg.StartLocation = p.Startlocation
	pkg.TripCategoryId = uint(p.CategoryId)

	err = c.Repo.CreatePackage(&pkg)
	if err != nil {
		return &cpb.Response{Status: "fail"}, errors.New("failed to create package")
	}
	return &cpb.Response{
		Status:  "success",
		Message: "package creation successful",
		Id:      int64(pkg.ID),
	}, nil
}

// AvailablePackageSvc retrieves available packages.
func (c *CoordinatorSVC) AvailablePackageSvc(p *cpb.View) (*cpb.PackagesResponse, error) {
	var pkgs []*cpb.Package
	var err error

	var packages []*dom.Package
	offset := 10 * (p.Page - 1)
	limit := 10

	if p.Status == "" {
		packages, err = c.Repo.FetchAllPackages(int(offset), limit)
	} else {
		packages, err = c.Repo.FetchPackages(int(offset), limit, p.Status)
	}

	if err != nil {
		return nil, errors.New("failed to fetch packages")
	}

	for _, pkge := range packages {
		pkg := &cpb.Package{
			PackageId:        int64(pkge.ID),
			Destination:      pkge.Destination,
			DestinationCount: int64(pkge.NumOfDestination),
			Enddate:          pkge.EndDate.Format("02-01-2006"),
			Image:            pkge.Images,
			Packagename:      pkge.Name,
			AvailableSpace:   int64(pkge.AvailableSpace),
			Price:            int64(pkge.MinPrice),
			Startdate:        pkge.StartDate.Format("02-01-2006"),
			Startlocation:    pkge.StartLocation,
			Description:      pkge.Description,
			MaxCapacity:      int64(pkge.MaxCapacity),
		}
		pkgs = append(pkgs, pkg)
	}

	return &cpb.PackagesResponse{Packages: pkgs}, nil
}

// ViewPackageSVC retrieves details of a specific package.
func (c *CoordinatorSVC) ViewPackageSVC(p *cpb.View) (*cpb.Package, error) {
	pkg, err := c.Repo.FetchPackage(uint(p.Id))
	if err != nil {
		return nil, errors.New("failed to fetch package")
	}

	ctgry := &cpb.Category{
		CategoryName: pkg.Category.Category,
	}

	destinations, err := c.Repo.FetchPackageDestination(pkg.ID)
	if err != nil {
		return nil, errors.New("failed to fetch package destination")
	}

	var dstn = []*cpb.Destination{}
	for _, dsn := range destinations {
		var ds = cpb.Destination{}
		ds.Description = dsn.Description
		ds.DestinationName = dsn.DestinationName
		ds.Image = dsn.Image
		ds.DestinationId = int64(dsn.Model.ID)
		ds.TransportationMode = dsn.TransportationMode
		ds.ArrivalLocation = dsn.ArrivalLocation
		dstn = append(dstn, &ds)
	}

	return &cpb.Package{
		Packagename:      pkg.Name,
		Startlocation:    pkg.StartLocation,
		Startdate:        pkg.StartDate.Format("02-01-2006"),
		Starttime:        pkg.StartTime,
		Enddate:          pkg.EndDate.Format("02-01-2006"),
		Price:            int64(pkg.MinPrice),
		Image:            pkg.Images,
		DestinationCount: int64(pkg.NumOfDestination),
		Destination:      pkg.Destination,
		PackageId:        int64(pkg.ID),
		AvailableSpace:   int64(pkg.AvailableSpace),
		Description:      pkg.Description,
		MaxCapacity:      int64(pkg.MaxCapacity),
		Category:         ctgry,
		Destinations:     dstn,
	}, nil
}

// AdminPackageStatusSvc updates package status.
func (c *CoordinatorSVC) AdminPackageStatusSvc(p *cpb.View) (*cpb.Response, error) {
	err := c.Repo.PackageStatusUpdate(uint(p.Id))
	if err != nil {
		return &cpb.Response{Status: "fail"}, errors.New("failed to update package status")
	}
	return &cpb.Response{Status: "success"}, nil
}

// FilterPackageSvc filters packages based on criteria.
func (c *CoordinatorSVC) FilterPackageSvc(p *cpb.Filter) (*cpb.PackagesResponse, error) {
	query := "SELECT * FROM packages WHERE 1 = 1"

	if p.Departurtime != "" {
		val := fmt.Sprintf(" AND start_time >= '%s'", p.Departurtime)
		query += val
	}
	if p.MinPrice > 0 {
		val := fmt.Sprintf(" AND min_price BETWEEN %d AND %d ", p.MinPrice, p.MaxPrice)
		query += val
	}

	if p.CategoryId != 0 {
		val := fmt.Sprintf(" AND trip_category_id = %d ", p.CategoryId)
		query += val
	}

	// Add ORDER BY clause
	if p.OrderBy != "" {
		query += " ORDER BY min_price " + p.OrderBy
	}

	// Add LIMIT and OFFSET clauses for paging
	offset := 10 * (p.Page - 1)
	limit := 10
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	var packages []*dom.Package
	db := c.Repo.GetDB()
	rows := db.Raw(query).Scan(&packages)
	if rows.Error != nil {
		return nil, errors.New("failed to filter packages")
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
			AvailableSpace:   int64(pkge.AvailableSpace),
			Price:            int64(pkge.MinPrice),
			Startdate:        pkge.StartDate.Format("02-01-2006"),
			Startlocation:    pkge.StartLocation,
			Description:      pkge.Description,
			MaxCapacity:      int64(pkge.MaxCapacity),
		}
		pkgs = append(pkgs, pkg)
	}

	return &cpb.PackagesResponse{Packages: pkgs}, nil
}

// ViewPackagesSvc retrieves packages of a coordinator.
func (c *CoordinatorSVC) ViewPackagesSvc(p *cpb.View) (*cpb.PackagesResponse, error) {
	offset := 10 * (p.Page - 1)
	limit := 10
	rslt, err := c.Repo.FindCoordinatorPackages(int(offset), limit, uint(p.Id))
	if err != nil {
		return nil, errors.New("failed to find packages")
	}

	var packages []*cpb.Package
	for _, pkges := range rslt {
		var pkg cpb.Package

		pkg.PackageId = int64(pkges.ID)
		pkg.Destination = pkges.Destination
		pkg.DestinationCount = int64(pkges.NumOfDestination)
		pkg.Enddate = pkges.EndDate.Format("02-01-2006")
		pkg.Image = pkges.Images
		pkg.Packagename = pkges.Name
		pkg.AvailableSpace = int64(pkges.AvailableSpace)
		pkg.Price = int64(pkges.MinPrice)
		pkg.Startdate = pkges.EndDate.Format("02-01-2006")
		pkg.Starttime = pkges.StartTime
		pkg.Startlocation = pkges.StartLocation
		pkg.Description = pkges.Description
		pkg.MaxCapacity = int64(pkges.MaxCapacity)

		packages = append(packages, &pkg)
	}

	return &cpb.PackagesResponse{Packages: packages}, nil
}
