package handler

import (
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorAddPackage(ctx context.Context, p *cpb.AddPackage) (*cpb.AddPackageResponce, error) {
	respnc, err := c.SVC.AddPackageSVC(p)
	if err != nil {
		log.Printf("Unable to create %v of email == %v, err: %v", p.Packagename, p.Coorinatoremail, err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorAddDestination(ctx context.Context, p *cpb.AddDestination) (*cpb.AddDestinationResponce, error) {
	respnc, err := c.SVC.AddDestinationSVC(p)
	if err != nil {
		log.Printf("Unable to create %v destination. err: %v", p.DestinationName, err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorAddActivity(ctx context.Context, p *cpb.AddActivity) (*cpb.AddActivityResponce, error) {
	respnc, err := c.SVC.AddActivitySVC(p)
	if err != nil {
		log.Printf("Unable to create %v activity. err: %v", p.ActivityName, err.Error())
		return nil, err
	}
	return respnc, nil
}
