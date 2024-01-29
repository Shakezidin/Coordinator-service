package handler

import (
	"context"
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

func (c *CoordinatorHandler) OnlinePayment(ctx context.Context, p *cpb.Booking) (*cpb.OnlinePaymentResponse, error) {
	respnc, err := c.SVC.OnlinePaymentSVC(ctx, p)
	if err != nil {
		log.Printf("error while adding payment details err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler)PaymentConfirmed(ctx context.Context,p *cpb.PaymentConfirmedRequest)(*cpb.BookingResponce,error){
	respnc, err := c.SVC.PaymentConfirmedSVC(ctx, p)
	if err != nil {
		log.Printf("error while adding payment details err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}