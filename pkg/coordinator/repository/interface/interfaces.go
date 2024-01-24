package interfaces

import (
	"time"

	cDOM "github.com/Shakezidin/pkg/entities/packages"
)

type CoordinatorRepoInter interface {
	FindUserByEmail(email string) (*cDOM.User, error)
	FindUserByPhone(number int) (*cDOM.User, error)
	CreateUser(user *cDOM.User) error
	FindCoordinatorPackages(id uint) (*[]cDOM.Package, error)
	CreatePackage(pkg *cDOM.Package) error
	CreateDestination(dtnt *cDOM.Destination) error
	CreateActivity(actvt *cDOM.Activity) error
	FetchPackages(offset, limit int, val string) ([]cDOM.Package, error)
	FetchPackage(id uint) (*cDOM.Package, error)
	FetchPackageDestination(id uint) ([]*cDOM.Destination, error)
	FecthDestination(id uint) (*cDOM.Destination, error)
	FecthDestinationActivity(id uint) ([]*cDOM.Activity, error)
	FecthActivity(id uint) (*cDOM.Activity, error)
	UpdatePassword(id uint, newpassword string) error
	CreateCatagory(catagory cDOM.Category) error
	FetchAllPackages(offset,limit int) (*[]cDOM.Package, error)
	PackageStatusUpdate(id uint) error
	FetchCatagories(offset,limit int) ([]*cDOM.Category, error) 
	FindUnboundedPackages(PickupPlace, Finaldestination string, MaxDestination int64,startDate, endDate time.Time) ([]*cDOM.Package, error)
	CreateTraveller(traveller cDOM.Traveller) error
	CreateActivityBooking(activity cDOM.ActivityBooking) error
	CreateBooking(booking cDOM.Booking) error
	FetchCatagory(catagory string) (*cDOM.Category, error)
	UpdatePackage(pkg *cDOM.Package)error
}
