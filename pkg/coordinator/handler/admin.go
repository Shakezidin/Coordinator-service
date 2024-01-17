package handler

import (
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) AddCategory(ctx context.Context, p *cpb.Category) (*cpb.Responce, error) {
	respnc, err := c.SVC.AddCatagorySVC(p)
	if err != nil {
		log.Printf("Unable to create %v catagory. err: %v", p.CategoryName, err.Error())
		return nil, err
	}
	return respnc, nil
}
