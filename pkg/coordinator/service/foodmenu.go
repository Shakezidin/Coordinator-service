package service

import (
	"errors"
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorSVC) AddFoodMenuSVC(p *cpb.FoodMenu) (*cpb.Responce, error) {
	_, err := c.Repo.FetchPackage(uint(p.PackageID))
	if err != nil {
		log.Print("package not found")
		return &cpb.Responce{
			Status:  "fail",
			Message: "package not found",
		}, errors.New("pacakge not found")
	}
	var foodmenu dom.FoodMenu

	foodmenu.PackageId = uint(p.PackageID)
	foodmenu.Breakfast = p.Breakfast
	foodmenu.Lunch = p.Lunch
	foodmenu.Dinner = p.Dinner
	foodmenu.Date = p.Date

	err = c.Repo.CreateFoodMenu(&foodmenu)
	if err != nil {
		return &cpb.Responce{
			Status:  "failure",
			Message: "destination creation error",
		}, err
	}
	return &cpb.Responce{
		Status:  "Success",
		Message: "foodmeny creation done",
		Id:      int64(foodmenu.ID),
	}, nil
}
