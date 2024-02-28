package handler

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorAddActivity(ctx context.Context, p *cpb.Activity) (*cpb.Response, error) {
	respnc, err := c.SVC.AddActivitySVC(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorViewActivity(ctx context.Context, p *cpb.View) (*cpb.Activity, error) {
	respnc, err := c.SVC.ViewActivitySvc(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}
