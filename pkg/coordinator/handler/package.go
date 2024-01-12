package handler

import (
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorAddPackage(ctx context.Context, p *cpb.AddPackage) (*cpb.AddPackageResponce, error) {
	respnc, err := c.SVC.AddPackageSVC(p)
	if err != nil {
		log.Printf("Unable to create %v of email == %v, err: %v", p.Packagename, p.Coorinatoremail, err.Error())
		return nil, err
	}
	return respnc, nil
}
