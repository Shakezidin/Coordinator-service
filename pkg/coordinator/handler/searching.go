package handler

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) PackageSearch(ctx context.Context, p *cpb.Search) (*cpb.PackagesResponse, error) {
	respnc, err := c.SVC.SearchPackageSVC(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}
