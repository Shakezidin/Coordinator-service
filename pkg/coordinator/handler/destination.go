package handler

import (
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorAddDestination(ctx context.Context, p *cpb.Destination) (*cpb.Responce, error) {
	respnc, err := c.SVC.AddDestinationSVC(p)
	if err != nil {
		log.Printf("Unable to create %v destination. err: %v", p.DestinationName, err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorViewDestination(ctx context.Context, p *cpb.View) (*cpb.Destination, error) {
	respnc, err := c.SVC.ViewDestinationSvc(p)
	if err != nil {
		log.Printf("Unable to fetch destination. err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}
