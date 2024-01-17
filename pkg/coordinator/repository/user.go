package repository

import cDOM "github.com/Shakezidin/pkg/entities/packages"

func (c *CoordinatorRepo) FindUserByEmail(email string) (*cDOM.User, error) {
	var user cDOM.User
	if err := c.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *CoordinatorRepo) FindUserByPhone(number int) (*cDOM.User, error) {
	var user cDOM.User
	if err := c.db.Where("phone = ?", number).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *CoordinatorRepo) CreateUser(user *cDOM.User) error {
	if err := c.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) UpdatePassword(id uint, newpassword string) error {
	user := cDOM.User{}
	user.ID = id

	result := c.db.Model(&user).Update("password", newpassword)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
