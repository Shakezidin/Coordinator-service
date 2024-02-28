package handler

import (
	"errors"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorSignupRequest(ctx context.Context, p *cpb.Signup) (*cpb.Response, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	respnc, err := c.SVC.SignupSVC(p)
	if err != nil {
		return respnc, err
	}

	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorSignupVerifyRequest(ctx context.Context, p *cpb.Verify) (*cpb.Response, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	respnc, err := c.SVC.VerifySVC(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorLoginRequest(ctx context.Context, p *cpb.Login) (*cpb.LoginResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	respnc, err := c.SVC.UserLogin(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}
