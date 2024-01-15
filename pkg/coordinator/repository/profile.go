package repository

import(
	cDOM "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorRepo) UpdatePassword(id uint, newpassword string) error {
	user := cDOM.User{}
	user.ID = id

	result := c.db.Model(&user).Update("password", newpassword)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
