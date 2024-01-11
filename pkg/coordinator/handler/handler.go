package handler

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	SVCinter "github.com/Shakezidin/pkg/coordinator/service/interface"
	"golang.org/x/net/context"
)

type CoordinatorHandler struct {
	SVC SVCinter.CoordinatorSVCInter
	cpb.CoordinatorServer
}

func (c *CoordinatorHandler) CoordinatorSignupRequest(ctx context.Context, p *cpb.Signup) (*cpb.SignupResponce, error) {
	result, err := c.SVC.SignupSVC(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewCoordinatorHandler(svc SVCinter.CoordinatorSVCInter) *CoordinatorHandler {
	return &CoordinatorHandler{
		SVC: svc,
	}
}
