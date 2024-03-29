package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

// SearchPackageSVC searches for packages based on search criteria.
func (c *CoordinatorSVC) SearchPackageSVC(p *cpb.Search) (*cpb.PackagesResponse, error) {
	if len(p.Destination) <= 1 {
		packages, err := UnboundedPackages(p.Pickup_Place, p.Final_Destination, p.Date, p.End_Date, p.Max_Destination, c, p.Page)
		if err != nil {
			return nil, err
		}
		return packages, nil
	}
	pkgs, err := BoundedPackages(c, p)
	if err != nil {
		return nil, err
	}
	return pkgs, nil
}

// UnboundedPackages retrieves unbounded packages.
func UnboundedPackages(PickupPlace, Finaldestination, date, enddate string, MaxDestination int64, svc *CoordinatorSVC, page int64) (*cpb.PackagesResponse, error) {
	startDate, err := time.Parse("02-01-2006", date)
	endDate, _ := time.Parse("02-01-2006", enddate)
	offset := 10 * (page - 1)
	limit := 10
	if err != nil {
		return nil, errors.New("error while parsing dates")
	}
	packages, err := svc.Repo.FindUnboundedPackages(int(offset), limit, PickupPlace, Finaldestination, MaxDestination, startDate, endDate)
	if err != nil {
		return nil, errors.New("error while finding packages")
	}

	var pkgs []*cpb.Package
	for _, pkges := range packages {
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

		pkgs = append(pkgs, &pkg)
	}
	return &cpb.PackagesResponse{Packages: pkgs}, nil
}

// BoundedPackages retrieves bounded packages.
func BoundedPackages(svc *CoordinatorSVC, p *cpb.Search) (*cpb.PackagesResponse, error) {
	startDate, err := time.Parse("02-01-2006", p.Date)
	endDate, _ := time.Parse("02-01-2006", p.End_Date)
	offset := 10 * (p.Page - 1)
	limit := 10
	if err != nil {
		return nil, errors.New("error while parsing dates")
	}

	packages, err := svc.Repo.FindUnboundedPackages(int(offset), limit, p.Pickup_Place, p.Final_Destination, p.Max_Destination, startDate, endDate)
	if err != nil {
		return nil, errors.New("error while finding packages")
	}

	var filteredPackages []*dom.Package

	for _, pkg := range packages {
		destinations, _ := svc.Repo.FetchPackageDestination(pkg.ID)
		var dsts []string
		for _, ds := range destinations {
			dsts = append(dsts, ds.DestinationName)
		}
		fmt.Println(p.Destination)
		if hasAllDestinations(dsts, p.Destination) {
			filteredPackages = append(filteredPackages, pkg)
		}
	}

	var pkgs []*cpb.Package
	for _, pkges := range filteredPackages {
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

		pkgs = append(pkgs, &pkg)
	}
	return &cpb.PackagesResponse{Packages: pkgs}, nil
}

// hasAllDestinations checks if a package has all specified destinations.
func hasAllDestinations(packageDestinations, searchDestinations []string) bool {
	for _, dest := range searchDestinations {
		if dest == "" {
			continue
		}

		found := false
		for _, pkgDest := range packageDestinations {

			if strings.EqualFold(strings.TrimSpace(dest), strings.TrimSpace(pkgDest)) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
