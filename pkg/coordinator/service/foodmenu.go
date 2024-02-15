package service

import (
	"errors"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

// AddFoodMenuSVC adds a food menu for a package.
func (c *CoordinatorSVC) AddFoodMenuSVC(p *cpb.FoodMenu) (*cpb.Response, error) {
	// Check if the package exists
	_, err := c.Repo.FetchPackage(uint(p.PackageID))
	if err != nil {
		return &cpb.Response{
			Status:  "fail",
			Message: "package not found",
		}, errors.New("package not found")
	}

	// Create the food menu
	var foodMenu dom.FoodMenu
	foodMenu.PackageID = uint(p.PackageID)
	foodMenu.Breakfast = p.Breakfast
	foodMenu.Lunch = p.Lunch
	foodMenu.Dinner = p.Dinner
	foodMenu.Date = p.Date

	err = c.Repo.CreateFoodMenu(&foodMenu)
	if err != nil {
		return &cpb.Response{
			Status:  "failure",
			Message: "food menu creation error",
		}, errors.New("error while creating food menu")
	}

	return &cpb.Response{
		Status: "Success",
		Id:     int64(foodMenu.ID),
	}, nil
}

// ViewFoodMenuSVC retrieves food menus for a package.
func (c *CoordinatorSVC) ViewFoodMenuSVC(p *cpb.View) (*cpb.FoodMenus, error) {
	// Calculate offset and limit for pagination
	offset := 10 * (p.Page - 1)
	limit := 10

	// Fetch food menus from the repository
	foodMenus, err := c.Repo.FetchFoodMenus(int(offset), limit, uint(p.Id))
	if err != nil {
		return nil, errors.New("error while fetching food menus")
	}

	// Convert domain food menus to protobuf food menus
	var pbFoodMenus []*cpb.FoodMenu
	for _, menu := range *foodMenus {
		pbFoodMenus = append(pbFoodMenus, &cpb.FoodMenu{
			FoodMenuId: int64(menu.ID),
			PackageID:  int64(menu.PackageID),
			Breakfast:  menu.Breakfast,
			Lunch:      menu.Lunch,
			Dinner:     menu.Dinner,
			Date:       menu.Date,
		})
	}

	return &cpb.FoodMenus{
		Foodmenu: pbFoodMenus,
	}, nil
}
