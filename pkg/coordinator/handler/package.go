package handler

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorAddPackage(ctx context.Context, p *cpb.Package) (*cpb.Response, error) {
	respnc, err := c.SVC.AddPackageSVC(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorViewPackage(ctx context.Context, p *cpb.View) (*cpb.Package, error) {
	respnc, err := c.SVC.ViewPackageSVC(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) AvailablePackages(ctx context.Context, p *cpb.View) (*cpb.PackagesResponse, error) {
	respnc, err := c.SVC.AvailablePackageSvc(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) AdminPacakgeStatus(ctx context.Context, p *cpb.View) (*cpb.Response, error) {
	respnc, err := c.SVC.AdminPackageStatusSvc(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) FilterPackage(ctx context.Context, p *cpb.Filter) (*cpb.PackagesResponse, error) {
	respnc, err := c.SVC.FilterPackageSvc(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorViewPackages(ctx context.Context, p *cpb.View) (*cpb.PackagesResponse, error) {
	respnc, err := c.SVC.ViewPackagesSvc(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}
