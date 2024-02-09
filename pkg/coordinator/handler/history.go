package handler

import (
	"context"
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

func (c *CoordinatorHandler) ViewHistory(ctx context.Context, p *cpb.View) (*cpb.Histories, error) {
	respnc, err := c.SVC.ViewhistorySVC(p)
	if err != nil {
		log.Printf("error while fetching history details err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) ViewBooking(ctx context.Context, p *cpb.View) (*cpb.History, error) {
	respnc, err := c.SVC.ViewBookingSVC(p)
	if err != nil {
		log.Printf("error while fetching history details err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CancelBooking(ctx context.Context, p *cpb.View) (*cpb.Responce, error) {
	respnc, err := c.SVC.CancelBookingSVC(p)
	if err != nil {
		log.Printf("error while fetching history details err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) ViewTraveller(ctx context.Context, p *cpb.View) (*cpb.TravellerDetails, error) {
	respnc, err := c.SVC.ViewTravellerSVC(p)
	if err != nil {
		log.Printf("error while fetching traveller details err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) SearchBooking(ctx context.Context, p *cpb.BookingSearchCriteria) (*cpb.Histories, error) {
	respnc, err := c.SVC.SearchBookingSVC(p)
	if err != nil {
		log.Printf("error while fetching traveller details err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}
