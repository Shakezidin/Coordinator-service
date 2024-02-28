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

	startDate, err := time.Parse(layout, p.Start_Date)
	if err != nil {
		return &cpb.Response{Status: "fail"}, errors.New("failed to parse startdate")
	}

	endDate, err := time.Parse(layout, p.End_Date)
	if err != nil {
		return &cpb.Response{Status: "fail"}, errors.New("failed to parse enddate")
	}

	pkg.CoordinatorID = uint(p.Coorinator_ID)
	pkg.Description = p.Description
	pkg.Destination = p.Destination
	pkg.EndDate = endDate
	pkg.Images = p.Image
	pkg.MaxCapacity = int(p.Max_Capacity)
	pkg.Name = p.Package_Name
	pkg.AvailableSpace = int(p.Max_Capacity)
	pkg.NumOfDestination = int(p.Destination_Count)
	pkg.MinPrice = int(p.Price)
	pkg.StartDate = startDate
	pkg.StartTime = p.Start_Time
	pkg.StartLocation = p.Start_Location
	pkg.TripCategoryId = uint(p.Category_ID)

	err = c.Repo.CreatePackage(&pkg)
	if err != nil {
		return &cpb.Response{Status: "fail"}, errors.New("failed to create package")
	}
	return &cpb.Response{
		Status:  "success",
		Message: "package creation successful",
		ID:      int64(pkg.ID),
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
			Package_ID:        int64(pkge.ID),
			Destination:       pkge.Destination,
			Destination_Count: int64(pkge.NumOfDestination),
			End_Date:          pkge.EndDate.Format("02-01-2006"),
			Image:             pkge.Images,
			Package_Name:      pkge.Name,
			Available_Space:   int64(pkge.AvailableSpace),
			Price:             int64(pkge.MinPrice),
			Start_Date:        pkge.StartDate.Format("02-01-2006"),
			Start_Location:    pkge.StartLocation,
			Description:       pkge.Description,
			Max_Capacity:      int64(pkge.MaxCapacity),
		}
		pkgs = append(pkgs, pkg)
	}

	return &cpb.PackagesResponse{Packages: pkgs}, nil
}

// ViewPackageSVC retrieves details of a specific package.
func (c *CoordinatorSVC) ViewPackageSVC(p *cpb.View) (*cpb.Package, error) {
	pkg, err := c.Repo.FetchPackage(uint(p.ID))
	if err != nil {
		return nil, errors.New("failed to fetch package")
	}

	ctgry := &cpb.Category{
		Category_Name: pkg.Category.Category,
	}

	destinations, err := c.Repo.FetchPackageDestination(pkg.ID)
	if err != nil {
		return nil, errors.New("failed to fetch package destination")
	}

	var dstn = []*cpb.Destination{}
	for _, dsn := range destinations {
		var ds = cpb.Destination{}
		ds.Description = dsn.Description
		ds.Destination_Name = dsn.DestinationName
		ds.Image = dsn.Image
		ds.Destination_ID = int64(dsn.Model.ID)
		ds.Transportation_Mode = dsn.TransportationMode
		ds.Arrival_Location = dsn.ArrivalLocation
		dstn = append(dstn, &ds)
	}

	return &cpb.Package{
		Package_Name:      pkg.Name,
		Start_Location:    pkg.StartLocation,
		Start_Date:        pkg.StartDate.Format("02-01-2006"),
		Start_Time:        pkg.StartTime,
		End_Date:          pkg.EndDate.Format("02-01-2006"),
		Price:             int64(pkg.MinPrice),
		Image:             pkg.Images,
		Destination_Count: int64(pkg.NumOfDestination),
		Destination:       pkg.Destination,
		Package_ID:        int64(pkg.ID),
		Available_Space:   int64(pkg.AvailableSpace),
		Description:       pkg.Description,
		Max_Capacity:      int64(pkg.MaxCapacity),
		Category:          ctgry,
		Destinations:      dstn,
	}, nil
}

// AdminPackageStatusSvc updates package status.
func (c *CoordinatorSVC) AdminPackageStatusSvc(p *cpb.View) (*cpb.Response, error) {
	err := c.Repo.PackageStatusUpdate(uint(p.ID))
	if err != nil {
		return &cpb.Response{Status: "fail"}, errors.New("failed to update package status")
	}
	return &cpb.Response{Status: "success"}, nil
}

// FilterPackageSvc filters packages based on criteria.
func (c *CoordinatorSVC) FilterPackageSvc(p *cpb.Filter) (*cpb.PackagesResponse, error) {
	query := "SELECT * FROM packages WHERE 1 = 1"

	if p.Departure_Time != "" {
		val := fmt.Sprintf(" AND start_time >= '%s'", p.Departure_Time)
		query += val
	}
	if p.Min_Price > 0 {
		val := fmt.Sprintf(" AND min_price BETWEEN %d AND %d ", p.Min_Price, p.Max_Price)
		query += val
	}

	if p.Category_ID != 0 {
		val := fmt.Sprintf(" AND trip_category_id = %d ", p.Category_ID)
		query += val
	}

	// Add ORDER BY clause
	if p.Order_By != "" {
		query += " ORDER BY min_price " + p.Order_By
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
			Package_ID:        int64(pkge.ID),
			Destination:       pkge.Destination,
			Destination_Count: int64(pkge.NumOfDestination),
			End_Date:          pkge.EndDate.Format("02-01-2006"),
			Image:             pkge.Images,
			Package_Name:      pkge.Name,
			Available_Space:   int64(pkge.AvailableSpace),
			Price:             int64(pkge.MinPrice),
			Start_Date:        pkge.StartDate.Format("02-01-2006"),
			Start_Location:    pkge.StartLocation,
			Description:       pkge.Description,
			Max_Capacity:      int64(pkge.MaxCapacity),
		}
		pkgs = append(pkgs, pkg)
	}

	return &cpb.PackagesResponse{Packages: pkgs}, nil
}

// ViewPackagesSvc retrieves packages of a coordinator.
func (c *CoordinatorSVC) ViewPackagesSvc(p *cpb.View) (*cpb.PackagesResponse, error) {
	offset := 10 * (p.Page - 1)
	limit := 10
	rslt, err := c.Repo.FindCoordinatorPackages(int(offset), limit, uint(p.ID))
	if err != nil {
		return nil, errors.New("failed to find packages")
	}

	var packages []*cpb.Package
	for _, pkges := range rslt {
		var pkg cpb.Package

		pkg.Package_ID = int64(pkges.ID)
		pkg.Destination = pkges.Destination
		pkg.Destination_Count = int64(pkges.NumOfDestination)
		pkg.End_Date = pkges.EndDate.Format("02-01-2006")
		pkg.Image = pkges.Images
		pkg.Package_Name = pkges.Name
		pkg.Available_Space = int64(pkges.AvailableSpace)
		pkg.Price = int64(pkges.MinPrice)
		pkg.Start_Date = pkges.EndDate.Format("02-01-2006")
		pkg.Start_Time = pkges.StartTime
		pkg.Start_Location = pkges.StartLocation
		pkg.Description = pkges.Description
		pkg.Max_Capacity = int64(pkges.MaxCapacity)

		packages = append(packages, &pkg)
	}

	return &cpb.PackagesResponse{Packages: packages}, nil
}
