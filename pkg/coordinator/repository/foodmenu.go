package repository

import cDOM "github.com/Shakezidin/pkg/entities/packages"

func (c *CoordinatorRepo) CreateFoodMenu(foodmenu *cDOM.FoodMenu) error {
	if err := c.DB.Create(&foodmenu).Error; err != nil {
		return err
	}
	return nil
}
