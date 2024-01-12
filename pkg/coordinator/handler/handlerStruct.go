package handler

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	SVCinter "github.com/Shakezidin/pkg/coordinator/service/interface"
)

type CoordinatorHandler struct {
	SVC SVCinter.CoordinatorSVCInter
	cpb.CoordinatorServer
}

func NewCoordinatorHandler(svc SVCinter.CoordinatorSVCInter) *CoordinatorHandler {
	return &CoordinatorHandler{
		SVC: svc,
	}
}
