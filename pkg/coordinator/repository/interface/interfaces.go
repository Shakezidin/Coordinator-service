package interfaces

import (
	cDOM "github.com/Shakezidin/pkg/DOM/coordinator"
)

type CoordinatorRepoInter interface {
	SignupRepo(user *cDOM.User) error
	FindUserByEmail(email string) (*cDOM.User, error)
	FindUserByPhone(number int)(*cDOM.User,error)
	CreateUser(user *cDOM.User) error
	FindCoordinatorPackages(id uint) (*[]cDOM.Package, error)
}
