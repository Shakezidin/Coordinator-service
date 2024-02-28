package handler

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) AddCategory(ctx context.Context, p *cpb.Category) (*cpb.Response, error) {
	respnc, err := c.SVC.AddCategorySVC(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler) Viewcategories(ctx context.Context, p *cpb.View) (*cpb.Categories, error) {
	respnc, err := c.SVC.ViewCategoriesSVC(p)
	if err != nil {
		return respnc, err
	}
	return respnc, nil
}
