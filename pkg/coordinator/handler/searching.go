package handler

import (
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) PackageSearch(ctx context.Context, p *cpb.Search) (*cpb.PackagesResponce, error) {
	respnc, err := c.SVC.SearchPackageSVC(p)
	if err != nil {
		log.Printf("Unable to fetch packages. err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}
