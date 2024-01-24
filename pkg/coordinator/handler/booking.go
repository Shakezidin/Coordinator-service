package handler

import (
	"context"
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

func (c *CoordinatorHandler) TravellerDetails(ctx context.Context, p *cpb.TravellerRequest) (*cpb.TravellerResponse, error) {
	respnc, err := c.SVC.TravellerDetails(p)
	if err != nil {
		log.Printf("error while adding traveller details err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) OfflineBooking(ctx context.Context, p *cpb.Booking) (*cpb.BookingResponce, error) {
	respnc, err := c.SVC.OfflineBooking(ctx,p)
	if err != nil {
		log.Printf("error while confirming bopking details err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}
