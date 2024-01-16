package interfaces

import (
	cDOM "github.com/Shakezidin/pkg/entities/packages"
)

type CoordinatorRepoInter interface {
	SignupRepo(user *cDOM.User) error
	FindUserByEmail(email string) (*cDOM.User, error)
	FindUserByPhone(number int) (*cDOM.User, error)
	CreateUser(user *cDOM.User) error
	FindCoordinatorPackages(id uint) (*[]cDOM.Package, error)
	CreatePackage(pkg *cDOM.Package) error
	CreateDestination(dtnt *cDOM.Destination) error
	CreateActivity(actvt *cDOM.Activity) error
	FetchAllPackages()(*[]cDOM.Package,error)
	FetchPackage(id uint)(*cDOM.Package,error)
	FetchPackageDestination(id uint)([]*cDOM.Destination,error)
	FecthDestination(id uint)(*cDOM.Destination,error)
	FecthDestinationActivity(id uint)([]*cDOM.Activity,error)
	FecthActivity(id uint) (*cDOM.Activity, error)
	UpdatePassword(id uint, newpassword string) error
}
