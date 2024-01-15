package interfaces

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

type CoordinatorSVCInter interface {
	SignupSVC(p *cpb.CoordinatorSignup) (*cpb.CoordinatorSignupResponce, error)
	VerifySVC(p *cpb.CoordinatorVerify) (*cpb.CoordinatorVerifyResponce, error)
	UserLogin(p *cpb.CoordinatorLogin) (*cpb.CoordinatorLoginResponce, error)
	AddPackageSVC(p *cpb.AddPackage) (*cpb.AddPackageResponce, error)
	AddDestinationSVC(p *cpb.AddDestination) (*cpb.AddDestinationResponce, error)
	AddActivitySVC(p *cpb.AddActivity) (*cpb.AddActivityResponce, error)
	AvailablePackageSvc() (*cpb.PackagesResponce, error)
	ViewPackageSVC(p *cpb.CoodinatorViewPackage)(*cpb.Package,error)
	ViewDestinationSvc(p *cpb.CoodinatorViewDestination) (*cpb.Destination, error)
	ViewActivitySvc(p *cpb.ViewActivity)(*cpb.Activity,error)
	NewPassword(p *cpb.Coordinatornewpassword) (*cpb.Coordinatornewpasswordresponce, error)
	ForgetPassword(p *cpb.CoordinatorforgetPassword) (*cpb.CoordinatorforgetPasswordResponce, error)
	ForgetPasswordVerify(p *cpb.CoordinatorforgetPasswordVerify) (*cpb.CoordinatorforgetPasswordVerifyResponce, error)
}
