package handler

import (
	"errors"
	"log"
	"time"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) ForgetPassword(ctx context.Context, p *cpb.CoordinatorforgetPassword) (*cpb.CoordinatorforgetPasswordResponce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.ForgetPassword(p)
	if err != nil {
		log.Printf("Unable to verify sent of otp for phone == %v, err: %v", p.Phone, err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *CoordinatorHandler) ForgetPasswordVerify(ctx context.Context, p *cpb.CoordinatorforgetPasswordVerify) (*cpb.CoordinatorforgetPasswordVerifyResponce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.ForgetPasswordVerify(p)
	if err != nil {
		log.Printf("Unable to verify otp for phone == %v, err: %v", p.Phone, err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *CoordinatorHandler) NewPassword(ctx context.Context, p *cpb.Coordinatornewpassword) (*cpb.Coordinatornewpasswordresponce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.NewPassword(p)
	if err != nil {
		log.Printf("Unable to update password, err: %v", err.Error())
		return nil, err
	}
	return resp, nil
}
