package repository

import cDOM "github.com/Shakezidin/pkg/entities/packages"

func (c *CoordinatorRepo) CreateFoodMenu(foodmenu *cDOM.FoodMenu) error {
	if err := c.DB.Create(&foodmenu).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) FetchFoodMenus(offset, limit int, id uint) (*[]cDOM.FoodMenu, error) {
	var foodmenu []cDOM.FoodMenu
	if err := c.DB.Offset(offset).Limit(limit).Where("package_id = ?", id).Find(&foodmenu).Error; err != nil {
		return nil, err
	}
	return &foodmenu, nil
}
