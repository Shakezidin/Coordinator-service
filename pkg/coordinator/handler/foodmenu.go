package handler

import (
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"golang.org/x/net/context"
)

func (c *CoordinatorHandler) CoordinatorAddFoodMenu(ctx context.Context, p *cpb.FoodMenu) (*cpb.Responce, error) {
	respnc, err := c.SVC.AddFoodMenuSVC(p)
	if err != nil {
		log.Printf("Unable to create foodmenu. err: %v", err.Error())
		return nil, err
	}
	return respnc, nil
}
