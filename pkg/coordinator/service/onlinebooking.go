package service

import (
	"context"
	"fmt"
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/razorpay/razorpay-go"
)

func (c *CoordinatorSVC) OnlinePaymentSVC(ctx context.Context, p *cpb.Booking) (*cpb.OnlinePaymentResponse, error) {

	UserId_Key := fmt.Sprintf("userId%d", p.RefId)
	userId, err := c.redis.Get(ctx, UserId_Key).Int()
	client := razorpay.NewClient(c.cfg.RAZORPAYKEYID, c.cfg.RAZORPAYSECRETKEY)

	amount_key := fmt.Sprintf("amount%d", p.RefId)

	totalFare, err := c.redis.Get(ctx, amount_key).Int()
	amountInPaise := int(totalFare) * 100

	data := map[string]interface{}{
		"amount":   amountInPaise,
		"currency": "INR",
		"receipt":  "bookingReference",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		log.Println("error creating order, err: ", err.Error())
		return nil, err
	}

	orderId := body["id"].(string)

	return &cpb.OnlinePaymentResponse{
		UserId:           int32(userId),
		TotalFare:        float32(totalFare),
		BookingReference: p.RefId,
		OrderId:          orderId,
	}, nil
}
