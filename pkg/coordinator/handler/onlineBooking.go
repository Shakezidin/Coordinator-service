package handler

import (
	"context"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

func (c *CoordinatorHandler) OnlinePayment(ctx context.Context, p *cpb.Booking) (*cpb.OnlinePaymentResponse, error) {
	respnc, err := c.SVC.OnlinePaymentSVC(ctx, p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) PaymentConfirmed(ctx context.Context, p *cpb.PaymentConfirmedRequest) (*cpb.BookingResponse, error) {
	respnc, err := c.SVC.PaymentConfirmedSVC(ctx, p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}
