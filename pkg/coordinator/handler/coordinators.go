package handler

import (
	"context"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

func (c *CoordinatorHandler) ViewCoordinators(ctx context.Context, p *cpb.View) (*cpb.Users, error) {
	respnc, err := c.SVC.ViewCoordinatorsSVC(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}
