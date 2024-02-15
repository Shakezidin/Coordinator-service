package handler

import (
	"context"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

func (c *CoordinatorHandler) UserTravellerDetails(ctx context.Context, p *cpb.TravellerRequest) (*cpb.TravellerResponse, error) {
	respnc, err := c.SVC.TravellerDetails(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}
