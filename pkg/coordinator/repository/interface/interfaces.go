package interfaces

import (
	"context"
	"time"

	cDOM "github.com/Shakezidin/pkg/entities/packages"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"gorm.io/gorm"
)

type CoordinatorRepoInter interface {
	FindUserByEmail(email string) (*cDOM.User, error)
	FindUserByPhone(number int) (*cDOM.User, error)
	CreateUser(user *cDOM.User) error
	FindCoordinatorPackages(offset, limit int, id uint) (*[]cDOM.Package, error)
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
	FetchAllPackages(offset, limit int) (*[]cDOM.Package, error)
	PackageStatusUpdate(id uint) error
	FetchCatagories(offset, limit int) ([]*cDOM.Category, error)
	FindUnboundedPackages(offset, limit int, PickupPlace, Finaldestination string, MaxDestination int64, startDate, endDate time.Time) ([]*cDOM.Package, error)
	CreateTraveller(traveller cDOM.Traveller) error
	CreateActivityBooking(activity cDOM.ActivityBooking) error
	CreateBooking(booking cDOM.Booking) error
	FetchCatagory(catagory string) (*cDOM.Category, error)
	UpdatePackage(pkg *cDOM.Package) error
	GetDB() *gorm.DB
	FetchUserById(id uint) (*cDOM.User, error)
	CreateFoodMenu(foodmenu *cDOM.FoodMenu) error
	FetchFoodMenus(offset, limit int, id uint) (*[]cDOM.FoodMenu, error)
	FetchHistory(offset, limit int, id uint) (*[]cDOM.Booking, error)
	FetchBooking(id uint) (*cDOM.Booking, error)
	UpdateBooking(booking cDOM.Booking) error
	UpdateUser(user *cDOM.User) error
	FetchBookings(offset, limit int, id uint) (*[]cDOM.Booking, error)
	FetchTraveller(id uint) (*cDOM.Traveller, error)
	CalculateDailyIncome(id uint, todayStart, todayEnd time.Time) int
	CalculateMonthlyIncome(id uint, currentMonthStart, currentMonthEnd time.Time) int
	AdminCalculateDailyIncome(todayStart, todayEnd time.Time) int
	AdminCalculateMonthlyIncome(currentMonthStart, currentMonthEnd time.Time) int
	FetchActivityBookingofUser(id uint) ([]*cDOM.ActivityBooking, error)
	FetchAllCoordinators(offset, limit int) (*[]cDOM.User, error)
	FetchNextDayTrip(date string) (*[]cDOM.Booking, error)
	UpdatePackageExpiration(date string) error
	CoordinatorCount() int64
	SearchBookings(ctx context.Context, criteria *cpb.BookingSearchCriteria) ([]*cDOM.Booking, error)
}
