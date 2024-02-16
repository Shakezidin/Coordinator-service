package interfaces

import (
	"context"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

type CoordinatorSVCInter interface {
	SignupSVC(p *cpb.Signup) (*cpb.Response, error)
	VerifySVC(p *cpb.Verify) (*cpb.Response, error)
	UserLogin(p *cpb.Login) (*cpb.LoginResponse, error)
	AddPackageSVC(p *cpb.Package) (*cpb.Response, error)
	AddDestinationSVC(p *cpb.Destination) (*cpb.Response, error)
	AddActivitySVC(p *cpb.Activity) (*cpb.Response, error)
	AvailablePackageSvc(p *cpb.View) (*cpb.PackagesResponse, error)
	ViewPackageSVC(p *cpb.View) (*cpb.Package, error)
	ViewDestinationSvc(p *cpb.View) (*cpb.Destination, error)
	ViewActivitySvc(p *cpb.View) (*cpb.Activity, error)
	NewPassword(p *cpb.NewPassword) (*cpb.Response, error)
	ForgetPassword(p *cpb.ForgetPassword) (*cpb.Response, error)
	ForgetPasswordVerify(p *cpb.ForgetPasswordVerify) (*cpb.Response, error)
	AddCategorySVC(p *cpb.Category) (*cpb.Response, error)
	ViewCategoriesSVC(p *cpb.View) (*cpb.Categories, error)
	AdminPackageStatusSvc(p *cpb.View) (*cpb.Response, error)
	SearchPackageSVC(p *cpb.Search) (*cpb.PackagesResponse, error)
	TravellerDetails(p *cpb.TravellerRequest) (*cpb.TravellerResponse, error)
	AddFoodMenuSVC(p *cpb.FoodMenu) (*cpb.Response, error)
	ViewFoodMenuSVC(p *cpb.View) (*cpb.FoodMenus, error)
	OnlinePaymentSVC(ctx context.Context, p *cpb.Booking) (*cpb.OnlinePaymentResponse, error)
	FilterPackageSvc(p *cpb.Filter) (*cpb.PackagesResponse, error)
	PaymentConfirmedSVC(ctx context.Context, p *cpb.PaymentConfirmedRequest) (*cpb.BookingResponse, error)
	ViewPackagesSvc(p *cpb.View) (*cpb.PackagesResponse, error)
	ViewHistorySVC(p *cpb.View) (*cpb.Histories, error)
	ViewBookingSVC(p *cpb.View) (*cpb.History, error)
	CancelBookingSVC(p *cpb.View) (*cpb.Response, error)
	ViewTravellerSVC(p *cpb.View) (*cpb.TravellerDetails, error)
	ViewDashBordSVC(p *cpb.View) (*cpb.Dashboard, error)
	ViewCoordinatorsSVC(p *cpb.View) (*cpb.Users, error)
	SearchBookingSVC(p *cpb.BookingSearchCriteria) (*cpb.Histories, error)
	FindCoordinatorByPackageId(pkgID uint) dom.User
	FetchNextDayTrip()
	UpdateExpiredPackage()
	StoreTravellerDetailsInRedis(ctx context.Context, refID string, p *cpb.TravellerRequest, pkg *dom.Package, travellers []dom.Traveller, activityBookings []dom.ActivityBooking, totalAmount int) error
	StoreInRedis(ctx context.Context, key string, data interface{}) error
	CalculateActivityTotal(travellerDetails []*cpb.TravellerDetails) int
}
