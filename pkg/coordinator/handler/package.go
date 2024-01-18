package handler

import (
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorAddPackage(ctx context.Context, p *cpb.Package) (*cpb.Responce, error) {
	respnc, err := c.SVC.AddPackageSVC(p)
	if err != nil {
		log.Printf("Unable to create %v ,err: %v", p.Packagename, err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) CoordinatorViewPackage(ctx context.Context, p *cpb.View) (*cpb.Package, error) {
	respnc, err := c.SVC.ViewPackageSVC(p)
	if err != nil {
		log.Printf("Unable to fetch package. err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) AvailablePackages(ctx context.Context, p *cpb.View) (*cpb.PackagesResponce, error) {
	respnc, err := c.SVC.AvailablePackageSvc(p)
	if err != nil {
		log.Printf("Unable to fetch packages. err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) AdminPacakgeStatus(ctx context.Context, p *cpb.View) (*cpb.Responce, error) {
	respnc, err := c.SVC.AdminPackageStatusSvc(p)
	if err != nil {
		log.Printf("Unable to fetch packages. err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}
