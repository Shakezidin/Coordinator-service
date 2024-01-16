package handler

import (
	"errors"
	"log"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorForgetPassword(ctx context.Context, p *cpb.ForgetPassword) (*cpb.Responce, error) {
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

func (c *CoordinatorHandler) CoordinatorForgetPasswordVerify(ctx context.Context, p *cpb.ForgetPasswordVerify) (*cpb.Responce, error) {
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

func (c *CoordinatorHandler) CoordinatorNewPassword(ctx context.Context, p *cpb.Newpassword) (*cpb.Responce, error) {
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
