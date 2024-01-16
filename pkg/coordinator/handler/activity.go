package handler

import (
	"log"

	"golang.org/x/net/context"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

func (c *CoordinatorHandler) CoordinatorAddActivity(ctx context.Context, p *cpb.Activity) (*cpb.Responce, error) {
	respnc, err := c.SVC.AddActivitySVC(p)
	if err != nil {
		log.Printf("Unable to create %v activity. err: %v", p.Activityname, err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorViewActivity(ctx context.Context, p *cpb.View) (*cpb.Activity, error) {
	respnc, err := c.SVC.ViewActivitySvc(p)
	if err != nil {
		log.Printf("Unable to fetch activity. err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}
