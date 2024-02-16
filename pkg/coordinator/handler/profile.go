package handler

import (
	"errors"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorForgetPassword(ctx context.Context, p *cpb.ForgetPassword) (*cpb.Response, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return &cpb.Response{
			Status: "fail",
		}, errors.New("deadline passed, aborting gRPC call")
	}
	respnc, err := c.SVC.ForgetPassword(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorForgetPasswordVerify(ctx context.Context, p *cpb.ForgetPasswordVerify) (*cpb.Response, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return &cpb.Response{
			Status: "fail",
		}, errors.New("deadline passed, aborting gRPC call")
	}

	respnc, err := c.SVC.ForgetPasswordVerify(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorNewPassword(ctx context.Context, p *cpb.Newpassword) (*cpb.Response, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return &cpb.Response{
			Status: "fail",
		}, errors.New("deadline passed, aborting gRPC call")
	}

	respnc, err := c.SVC.NewPassword(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) ViewDashboard(ctx context.Context, p *cpb.View) (*cpb.Dashboard, error) {
	respnc, err := c.SVC.ViewDashBordSVC(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}
