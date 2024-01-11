package interfaces

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

type CoordinatorSVCInter interface {
	SignupSVC(p *cpb.Signup) (*cpb.SignupResponce,error)
}
