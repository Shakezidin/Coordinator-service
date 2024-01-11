package service

import (
	"errors"

	cDOM "github.com/Shakezidin/pkg/DOM/coordinator"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	inter "github.com/Shakezidin/pkg/coordinator/repository/interface"
	SVCinter "github.com/Shakezidin/pkg/coordinator/service/interface"
)

type CoordinatorSVC struct {
	Repo inter.CoordinatorRepoInter
}

func (c *CoordinatorSVC) SignupSVC(p *cpb.Signup) (*cpb.SignupResponce,error) {
	var user cDOM.User

	user.Name = p.Name
	user.Password = p.Password
	user.Phone = int(p.Phone)
	user.Email = p.Email
	user.Role = "coordinator"

	if err := c.Repo.SignupRepo(&user); err != nil {
		return nil,errors.New("user creation error")
	}

	return &cpb.SignupResponce{
		Status: "success",
		Message: "Coordinator signup done",
	},nil
}

func NewCoordinatorSVC(repo inter.CoordinatorRepoInter) SVCinter.CoordinatorSVCInter {
	return &CoordinatorSVC{
		Repo: repo,
	}
}
