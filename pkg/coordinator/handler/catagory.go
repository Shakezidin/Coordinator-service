package handler

import (
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) AddCatagory(ctx context.Context, p *cpb.Category) (*cpb.Responce, error) {
	respnc, err := c.SVC.AddCatagorySVC(p)
	if err != nil {
		log.Printf("Unable to create %v catagory. err: %v", p.CategoryName, err.Error())
		return nil, err
	}
	return respnc, nil
}

func (c *CoordinatorHandler)ViewCatagories(ctx context.Context,p *cpb.View)(*cpb.Catagories,error){
	respnc, err := c.SVC.ViewCatagoriesSVC(p)
	if err != nil {
		log.Printf("Unable to fetch catagories. err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}