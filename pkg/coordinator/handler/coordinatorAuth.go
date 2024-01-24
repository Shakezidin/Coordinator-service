package handler

import (
	"errors"
	"log"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorSignupRequest(ctx context.Context, p *cpb.Signup) (*cpb.Responce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	result, err := c.SVC.SignupSVC(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CoordinatorHandler) CoordinatorSignupVerifyRequest(ctx context.Context, p *cpb.Verify) (*cpb.Responce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.VerifySVC(p)
	if err != nil {
		log.Printf("Unable to verify %v of email == %v, err: %v", p.OTP, p.Email, err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *CoordinatorHandler) CoordinatorLoginRequest(ctx context.Context, p *cpb.Login) (*cpb.LoginResponce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	respnc, err := c.SVC.UserLogin(p)
	if err != nil {
		log.Printf("Unable to login %v of email == %v, err: %v", p.Role, p.Email, err.Error())
		return nil, err
	}
	return respnc, nil
}
