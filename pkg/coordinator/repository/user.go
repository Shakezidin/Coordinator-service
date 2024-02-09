package repository

import (
	"time"

	cDOM "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorRepo) FindUserByEmail(email string) (*cDOM.User, error) {
	var user cDOM.User
	if err := c.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *CoordinatorRepo) FindUserByPhone(number int) (*cDOM.User, error) {
	var user cDOM.User
	if err := c.DB.Where("phone = ?", number).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *CoordinatorRepo) CreateUser(user *cDOM.User) error {
	if err := c.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) UpdatePassword(id uint, newpassword string) error {
	user := cDOM.User{}
	user.ID = id

	result := c.DB.Model(&user).Update("password", newpassword)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *CoordinatorRepo) FetchUserById(id uint) (*cDOM.User, error) {
	var user cDOM.User
	if err := c.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *CoordinatorRepo) UpdateUser(user *cDOM.User) error {
	if err := c.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) CalculateDailyIncome(id uint, todayStart, todayEnd time.Time) int {
	var todayIncome int
	query := `SELECT COALESCE(SUM(paid_price), 0) FROM "bookings" WHERE (coordinator_id = ? AND book_date >= ? AND book_date <= ? AND cancelled_status = ? AND payment_mode = ?) AND "bookings"."deleted_at" IS NULL`
	c.DB.Raw(query, id, todayStart, todayEnd, "false", "full amount").Scan(&todayIncome)
	return todayIncome
}

func (c *CoordinatorRepo) CalculateMonthlyIncome(id uint, currentMonthStart, currentMonthEnd time.Time) int {
	var monthlyIncome int
	query := `SELECT COALESCE(SUM(paid_price), 0) FROM "bookings" WHERE (coordinator_id = ? AND book_date >= ? AND book_date <= ? AND cancelled_status = ? AND payment_mode = ?) AND "bookings"."deleted_at" IS NULL`
	c.DB.Raw(query, id, currentMonthStart, currentMonthEnd, "false", "full amount").Scan(&monthlyIncome)
	return monthlyIncome
}

func (c *CoordinatorRepo) AdminCalculateDailyIncome(todayStart, todayEnd time.Time) int {
	var todayIncome int
	query := `SELECT COALESCE(SUM(package_price), 0) FROM "bookings" WHERE (book_date >= ? AND book_date <= ? AND cancelled_status = ?) AND "bookings"."deleted_at" IS NULL`
	c.DB.Raw(query, todayStart, todayEnd, "false").Scan(&todayIncome)
	return todayIncome
}

func (c *CoordinatorRepo) AdminCalculateMonthlyIncome(currentMonthStart, currentMonthEnd time.Time) int {
	var todayIncome int
	query := `SELECT COALESCE(SUM(package_price), 0) FROM "bookings" WHERE (book_date >= ? AND book_date <= ? AND cancelled_status = ?) AND "bookings"."deleted_at" IS NULL`
	c.DB.Raw(query, currentMonthStart, currentMonthEnd, "false").Scan(&todayIncome)
	return todayIncome
}

func (c *CoordinatorRepo) FetchAllCoordinators(offset, limit int) (*[]cDOM.User, error) {
	var coordinator *[]cDOM.User
	if err := c.DB.Offset(offset).Limit(limit).Find(&coordinator).Error; err != nil {
		return nil, err
	}
	return coordinator, nil
}

func (c *CoordinatorRepo) CoordinatorCount() int64 {
	var coordinatorCount int64
	c.DB.Model(&cDOM.User{}).Count(&coordinatorCount)
	return coordinatorCount
}
