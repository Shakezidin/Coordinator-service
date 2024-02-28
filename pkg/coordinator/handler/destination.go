package handler

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorAddDestination(ctx context.Context, p *cpb.Destination) (*cpb.Response, error) {
	respnc, err := c.SVC.AddDestinationSVC(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorViewDestination(ctx context.Context, p *cpb.View) (*cpb.Destination, error) {
	respnc, err := c.SVC.ViewDestinationSvc(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}
