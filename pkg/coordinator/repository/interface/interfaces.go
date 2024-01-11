package interfaces

import (
	cDOM "github.com/Shakezidin/pkg/DOM/coordinator"
)

type CoordinatorRepoInter interface {
	SignupRepo(user *cDOM.User) error
}
