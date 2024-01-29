package interfaces

import (
	"context"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

type CoordinatorSVCInter interface {
	SignupSVC(p *cpb.Signup) (*cpb.Responce, error)
	VerifySVC(p *cpb.Verify) (*cpb.Responce, error)
	UserLogin(p *cpb.Login) (*cpb.LoginResponce, error)
	AddPackageSVC(p *cpb.Package) (*cpb.Responce, error)
	AddDestinationSVC(p *cpb.Destination) (*cpb.Responce, error)
	AddActivitySVC(p *cpb.Activity) (*cpb.Responce, error)
	AvailablePackageSvc(p *cpb.View) (*cpb.PackagesResponce, error)
	ViewPackageSVC(p *cpb.View) (*cpb.Package, error)
	ViewDestinationSvc(p *cpb.View) (*cpb.Destination, error)
	ViewActivitySvc(p *cpb.View) (*cpb.Activity, error)
	NewPassword(p *cpb.Newpassword) (*cpb.Responce, error)
	ForgetPassword(p *cpb.ForgetPassword) (*cpb.Responce, error)
	ForgetPasswordVerify(p *cpb.ForgetPasswordVerify) (*cpb.Responce, error)
	AddCatagorySVC(p *cpb.Category) (*cpb.Responce, error)
	ViewCatagoriesSVC(p *cpb.View) (*cpb.Catagories, error)
	AdminPackageStatusSvc(p *cpb.View) (*cpb.Responce, error)
	SearchPackageSVC(p *cpb.Search) (*cpb.PackagesResponce, error)
	TravellerDetails(p *cpb.TravellerRequest) (*cpb.TravellerResponse, error)
	// OfflineBooking(ctx context.Context, p *cpb.Booking) (*cpb.BookingResponce, error)
	OnlinePaymentSVC(ctx context.Context,p *cpb.Booking)(*cpb.OnlinePaymentResponse,error)
	FilterPackageSvc(p *cpb.Filter)(*cpb.PackagesResponce,error)
	PaymentConfirmedSVC(ctx context.Context,p *cpb.PaymentConfirmedRequest)(*cpb.BookingResponce,error)
}
